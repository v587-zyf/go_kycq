package godEquip

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

/**
 *  @Description: 神兵血炼
 *  @param user
 *  @param heroIndex
 *  @param godEquipId
 *  @param op
 *  @return error
 */
func (this *GodEquipManager) GodEquipBlood(user *objs.User, heroIndex, godEquipId int, op *ophelper.OpBagHelperDefault) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	godEquip, ok := hero.GodEquips[godEquipId]
	nextLv, nowLv := 1, 1
	if ok {
		if godEquip.Blood > 0 {
			nextLv = godEquip.Blood + 1
			nowLv = godEquip.Blood
		}
	} else {
		hero.GodEquips[godEquipId] = &model.GodEquip{Id: godEquipId}
		godEquip = hero.GodEquips[godEquipId]
	}
	if gamedb.GetGodBloodGodBloodCfg(gamedb.GetRealId(godEquipId, nextLv)) == nil {
		return gamedb.ERRLVENOUGH
	}
	bloodCfg := gamedb.GetGodBloodGodBloodCfg(gamedb.GetRealId(godEquipId, nowLv))
	if bloodCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("godBlood id:%v lv:%v", godEquipId, godEquip.Blood)
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, bloodCfg.Consume); err != nil {
		return err
	}

	godEquip.Blood = nextLv

	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}
