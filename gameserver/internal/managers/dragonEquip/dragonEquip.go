package dragonEquip

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewDragonEquipManager(module managersI.IModule) *DragonEquipManager {
	return &DragonEquipManager{IModule: module}
}

type DragonEquipManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 龙器升级
 *  @param user
 *  @param heroIndex
 *  @param id	龙器id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *DragonEquipManager) UpLv(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault, ack *pb.DragonEquipUpLvAck) error {
	if heroIndex <= 0 || id <= 0 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}

	heroDragonEquip := hero.DragonEquip
	lv, ok := heroDragonEquip[id]
	if !ok {
		lv = 0
	}
	if gamedb.GetDragonEquipLevelDragonEquipLevelCfg(gamedb.GetRealId(id, lv+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	dragonEquipLevelCfg := gamedb.GetDragonEquipLevelDragonEquipLevelCfg(gamedb.GetRealId(id, lv))
	if !this.GetCondition().CheckMulti(user, heroIndex, dragonEquipLevelCfg.Condition) {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, dragonEquipLevelCfg.Item); err != nil {
		return err
	}
	heroDragonEquip[id] = lv + 1

	ack.HeroIndex = int32(heroIndex)
	ack.Id = int32(id)
	ack.Lv = int32(heroDragonEquip[id])
	kyEvent.DragonEquipLvUp(user, id, heroDragonEquip[id])

	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetTask().UpdateTaskProcess(user, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HERO_DRAGON_EQUIP_GRADE, []int{})
	return nil
}
