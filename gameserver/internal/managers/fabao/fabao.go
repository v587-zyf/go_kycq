package fabao

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewFabaoManager(module managersI.IModule) *FabaoManager {
	return &FabaoManager{IModule: module}
}

type FabaoManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *FabaoManager) ActiveChange(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.FabaoActiveAck) error {
	_, err := this.CheckFabao(user, id, true)
	if err != nil {
		return err
	}

	fabaoConf := gamedb.GetFabaoFabaoCfg(id)
	itemId, itemCount := fabaoConf.ActiveItem, 1
	for _, v := range fabaoConf.Condition {
		if _, b := this.GetCondition().Check(user, -1, v.K, v.V); !b {
			return gamedb.ERRCONDITION
		}
	}
	err = this.GetBag().Remove(user, op, itemId, itemCount)
	if err != nil {
		return err
	}
	user.Fabaos[id] = &model.Fabao{
		Id:    id,
		Skill: make([]int, 0),
	}

	ack.Fabao = builder.BuilderFabaoUnit(user.Fabaos[id])
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

func (this *FabaoManager) UpLevel(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.FabaoUpLevelAck) error {
	fabao, err := this.CheckFabao(user, id, false)
	if err != nil {
		return err
	}

	curLevel, curExp := fabao.Level, fabao.Exp
	maxLv := gamedb.GetMaxValById(id, constMax.MAX_FABAO_LEVEL)
	if curLevel >= maxLv {
		return gamedb.ERRLVENOUGH
	}
	fabaoConf := gamedb.GetFabaoById(id)
	needItemId := fabaoConf.UpLvCostItem

	fabaoLvConf := gamedb.GetFabaoLvByIdAndLv(id, curLevel)
	effectVal := gamedb.GetItemBaseCfg(needItemId).EffectVal
	needExp := fabaoLvConf.Exp - curExp
	needItemCount := needExp / effectVal
	hasItemCount, _ := this.GetBag().GetItemNum(user, needItemId)
	beforeLv := fabao.Level
	if hasItemCount >= needItemCount {
		this.GetBag().Remove(user, op, needItemId, needItemCount)
		fabao.Level += 1
		fabao.Exp = 0
	} else {
		this.GetBag().Remove(user, op, needItemId, hasItemCount)
		fabao.Exp += hasItemCount * effectVal
	}

	kyEvent.FaBaoUp(user, id, beforeLv, fabao.Level)
	ack.Fabao = builder.BuilderFabaoUnit(fabao)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

func (this *FabaoManager) ActiveSkillChange(user *objs.User, id, skillId int, op *ophelper.OpBagHelperDefault, ack *pb.FabaoSkillActiveAck) error {
	fabao, err := this.CheckFabao(user, id, false)
	if err != nil {
		return err
	}

	for _, v := range fabao.Skill {
		if v == skillId {
			return gamedb.ERRFABAOSKILLREPEATACTIVE
		}
	}
	fabaoSkillConf := gamedb.GetFabaoSkillByIdAndSkillId(id, skillId)
	itemId, itemCount := fabaoSkillConf.SkillCostItem.ItemId, fabaoSkillConf.SkillCostItem.Count
	if fabao.Level < fabaoSkillConf.Fabao_level {
		return gamedb.ERRFABAOLVNOTENOUGH
	}
	err = this.GetBag().Remove(user, op, itemId, itemCount)
	if err != nil {
		return err
	}
	fabao.Skill = append(fabao.Skill, skillId)

	ack.Fabao = builder.BuilderFabaoUnit(fabao)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

func (this *FabaoManager) CheckFabao(user *objs.User, id int, checkActive bool) (fabao *model.Fabao, err error) {
	if id <= 0 {
		err = gamedb.ERRPARAM
		return
	}
	fabao = user.Fabaos[id]
	if checkActive {
		if fabao != nil {
			err = gamedb.ERRFABAOREPEATACTIVE
			return
		}
	} else {
		if fabao == nil {
			err = gamedb.ERRFABAONOTACTIVE
			return
		}
	}
	return
}
