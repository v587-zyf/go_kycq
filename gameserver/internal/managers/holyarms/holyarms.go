package holyarms

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewHolyarmsManager(module managersI.IModule) *HolyarmsManager {
	return &HolyarmsManager{IModule: module}
}

type HolyarmsManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 至尊法器激活
 *  @param user
 *  @param id	至尊法器id
 *  @param ack
 *  @return error
 */
func (this *HolyarmsManager) Active(user *objs.User, id int, ack *pb.HolyActiveAck) error {
	//holy := user.Holyarms[id]
	//if holy == nil {
	//	return gamedb.ERRPARAM
	//}
	//if holy.Level > 0 {
	//	return gamedb.ERRREPEATACTIVE
	//}
	//
	//holyConf := gamedb.GetHolyArmsHolyArmsCfg(id)
	//if check := this.GetCondition().CheckMulti(user, -1, holyConf.Condition); !check {
	//	return gamedb.ERRCONDITION
	//}
	//lvCfg := gamedb.GetHolyLvByHidAndLv(id, holy.Level)
	//if holy.Exp < lvCfg.Exp {
	//	return gamedb.ERRPARAM
	//}
	//
	//holy.Level++
	//holy.Exp -= lvCfg.Exp
	//user.Dirty = true
	//
	//ack.Holy = builder.BuilderHoly(id, holy)
	//
	//this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 至尊法器升级
 *  @param user
 *  @param id		至尊法器id
 *  @param itemId	增加经验道具id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *HolyarmsManager) UpLevel(user *objs.User, id, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.HolyUpLevelAck) error {
	if id < 1 || itemId < 1 {
		return gamedb.ERRPARAM
	}
	holyarm, ok := user.Holyarms[id]
	if !ok {
		user.Holyarms[id] = &model.Holyarm{Skill: make(model.IntKv)}
		holyarm = user.Holyarms[id]
	}
	curLv, curExp := holyarm.Level, holyarm.Exp
	maxLv := gamedb.GetMaxValById(id, constMax.MAX_HOLY_LEVEL)
	if curLv >= maxLv {
		return gamedb.ERRLVENOUGH
	}
	if err := this.GetBag().Remove(user, op, itemId, 1); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	curExp += gamedb.GetItemBaseCfg(itemId).EffectVal
	for i := curLv; i < maxLv; i++ {
		lvConf := gamedb.GetHolyLvByHidAndLv(id, i)
		if curExp >= lvConf.Exp {
			curLv++
			curExp -= lvConf.Exp
		}
	}
	holyarm.Level = curLv
	holyarm.Exp = curExp
	ack.Holy = builder.BuilderHoly(id, holyarm)
	this.GetUserManager().UpdateCombat(user, -1)
	this.GetTask().UpdateTaskProcess(user, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_HOLYARMS_UP_GRADE, []int{})
	return nil
}

/**
 *  @Description: 至尊法器一键升级（后期开放）
 *  @param user
 *  @param id	  至尊法器id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *HolyarmsManager) AutoUpLv(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.HolyUpLevelAck) error {
	if id < 1 {
		return gamedb.ERRPARAM
	}
	holy := user.Holyarms[id]
	if holy == nil {
		user.Holyarms[id] = &model.Holyarm{}
		holy = user.Holyarms[id]
	}

	maxLv := gamedb.GetMaxValById(id, constMax.MAX_HOLY_LEVEL)
	if holy.Level >= maxLv {
		return gamedb.ERRLVENOUGH
	}
	holyCfg := gamedb.GetHolyArmsHolyArmsCfg(id)
	if holyCfg == nil {
		return gamedb.ERRPARAM
	}

	needCostItemId := holyCfg.UpLvCostItem[0]
	itemCfg := gamedb.GetItemBaseCfg(needCostItemId)
	holyLvConf := gamedb.GetHolyLvByHidAndLv(id, holy.Level)
	needExp := holyLvConf.Exp - holy.Exp
	needItemNum := common.CeilFloat64(float64(needExp) / float64(itemCfg.EffectVal))
	if hasNum, _ := this.GetBag().GetItemNum(user, needCostItemId); hasNum >= needItemNum {
		if err := this.GetBag().Remove(user, op, needCostItemId, needItemNum); err != nil {
			return err
		}
		holy.Level++
		holy.Exp = itemCfg.EffectVal*needItemNum + holy.Exp - holyLvConf.Exp

		nextCfg := gamedb.GetHolyLvByHidAndLv(id, holy.Level)
		for holy.Exp >= nextCfg.Exp {
			holy.Exp -= nextCfg.Exp
			holy.Level++
			if nextCfg = gamedb.GetHolyLvByHidAndLv(id, holy.Level); nextCfg == nil {
				break
			}
		}
	} else {
		if err := this.GetBag().Remove(user, op, needCostItemId, hasNum); err != nil {
			return err
		}
		holy.Exp += itemCfg.EffectVal * hasNum
	}

	ack.Holy = builder.BuilderHoly(id, user.Holyarms[id])
	this.GetUserManager().UpdateCombat(user, -1)
	this.GetTask().UpdateTaskProcess(user, false, false)
	return nil
}

/**
 *  @Description: 至尊法器激活技能
 *  @param user
 *  @param hid	至尊法器id
 *  @param hlv	至尊法器等级
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *HolyarmsManager) ActiveSkill(user *objs.User, hid, hlv int, op *ophelper.OpBagHelperDefault, ack *pb.HolySkillActiveAck) error {
	holy := user.Holyarms[hid]
	if holy == nil {
		return gamedb.ERRPARAM
	}
	holySkill := holy.Skill
	if holySkill == nil {
		holy.Skill = make(model.IntKv)
		holySkill = holy.Skill
	}
	if holySkill[hlv] != 0 {
		return gamedb.ERRREPEATACTIVE
	}

	skillConf := gamedb.GetHolySkillByHidAndLv(hid, hlv, 0)
	if skillConf == nil {
		return gamedb.ERRPARAM
	}

	if err := this.GetBag().Remove(user, op, skillConf.SkillCostItem.ItemId, skillConf.SkillCostItem.Count); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	holy.Skill[hlv] = 1

	ack.Holy = builder.BuilderHoly(hid, holy)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 至尊法器技能升级
 *  @param user
 *  @param hid	至尊法器id
 *  @param hlv	至尊法器等级
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *HolyarmsManager) SkillUpLv(user *objs.User, hid, hlv int, op *ophelper.OpBagHelperDefault, ack *pb.HolySkillUpLvAck) error {
	if hid < 1 || hlv < 1 {
		return gamedb.ERRPARAM
	}
	holy := user.Holyarms[hid]
	if holy == nil {
		return gamedb.ERRPARAM
	}

	holySkill := holy.Skill
	if holySkill == nil {
		holy.Skill = make(model.IntKv)
		holySkill = holy.Skill
	}
	holySkillLv := holySkill[hlv]

	maxLv := gamedb.GetHolySkillMaxLv(hid, hlv)
	if holySkillLv >= maxLv {
		return gamedb.ERRLVENOUGH
	}

	skillConf := gamedb.GetHolySkillByHidAndLv(hid, hlv, holySkillLv)
	if skillConf == nil {
		return gamedb.ERRPARAM
	}
	if err := this.GetBag().Remove(user, op, skillConf.SkillCostItem.ItemId, skillConf.SkillCostItem.Count); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	holy.Skill[hlv]++

	ack.Holy = builder.BuilderHoly(hid, holy)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}
