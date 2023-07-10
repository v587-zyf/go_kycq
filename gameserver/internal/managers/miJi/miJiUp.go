package miJi

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type MiJi struct {
	util.DefaultModule
	managersI.IModule
}

func NewMiJi(module managersI.IModule) *MiJi {
	m := &MiJi{IModule: module}
	return m
}

func (this *MiJi) GetMiJiInfos(user *objs.User) []*pb.MiJiInfo {

	data := make([]*pb.MiJiInfo, 0)
	if len(user.MiJi) > 0 {
		for id, info := range user.MiJi {
			data = append(data, &pb.MiJiInfo{Id: int32(id), Lv: int32(info.MiJiLv)})
		}
	}
	return data
}

func (this *MiJi) Up(user *objs.User, id int, ack *pb.MiJiUpAck, op *ophelper.OpBagHelperDefault) error {

	cfg := gamedb.GetMijiMijiCfg(id)
	if cfg == nil {
		return gamedb.ERRPARAM
	}

	if user.MiJi[id] == nil {
		user.MiJi[id] = &model.MiJiUnit{}
	}
	lvCfgId := gamedb.GetRealId(id, user.MiJi[id].MiJiLv)
	mjCfg := gamedb.GetMijiLevelMijiLevelCfg(lvCfgId)
	if mjCfg == nil {
		return gamedb.ERRPARAM
	}
	if mjCfg.Consume == nil {
		return gamedb.ERRMAXLV
	}

	has, _ := this.GetBag().HasEnoughItems(user, mjCfg.Consume)
	if !has {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err := this.GetBag().RemoveItemsInfos(user, op, mjCfg.Consume)
	if err != nil {
		return err
	}

	user.MiJi[id].MiJiLv += 1
	user.MiJi[id].MiJiType = cfg.Type
	user.Dirty = true
	ack.Id = int32(id)
	ack.Lv = int32(user.MiJi[id].MiJiLv)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

func (this *MiJi) GetMiJiLvId(id, lv int) int {
	return lv + id*10000
}

func (this *MiJi) GetMiJiSkillInfo(user *objs.User) *pbserver.MiJiInfo {

	allSkillInfo := &pbserver.MiJiInfo{}
	allSkillInfo.Skills = make([]*pbserver.Skill, 0)
	for miJiId, data := range user.MiJi {
		lvId := this.GetMiJiLvId(miJiId, data.MiJiLv)
		lvCfg := gamedb.GetMijiLevelMijiLevelCfg(lvId)
		if lvCfg == nil {
			continue
		}
		skillInfo := gamedb.GetSkillLevelSkillCfg(lvCfg.SkillLevel)
		if skillInfo == nil {
			continue
		}
		allSkillInfo.Skills = append(allSkillInfo.Skills, &pbserver.Skill{Id: int32(lvCfg.SkillLevel), Level: int32(skillInfo.Level)})
	}
	return allSkillInfo
}
