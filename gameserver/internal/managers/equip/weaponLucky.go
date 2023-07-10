package equip

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 祝福油使用校验
 *  @param user
 *  @param heroIndex
 *  @param itemId
 *  @return error
 */
func (this *EquipManager) CheckWeaponLucky(user *objs.User, heroIndex, itemId int) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	equip := hero.Equips[pb.EQUIPTYPE_WEAPON_R]
	if equip.ItemId <= 0 {
		return gamedb.ERRNOTWEAREQUIP
	}
	maxLv := gamedb.GetMaxValById(0, constMax.MAX_BLESS_LEVEL)
	if equip.Lucky >= maxLv {
		return gamedb.ERRLVENOUGH
	}
	blessCfg := gamedb.GetBlessBlessCfg(equip.Lucky)
	if itemId != blessCfg.Consume.ItemId {
		return gamedb.ERRGOODS
	}
	itemNum, _ := this.GetBag().GetItemNum(user, blessCfg.Consume.ItemId)
	if itemNum <= 0 {
		return gamedb.ERRNOTENOUGHGOODS
	}
	return nil
}

func (this *EquipManager) WeaponBless(user *objs.User, heroIndex int) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}

	equip := user.Heros[heroIndex].Equips[pb.EQUIPTYPE_WEAPON_R]
	blessCfg := gamedb.GetBlessBlessCfg(equip.Lucky)
	equipBlessNtf := &pb.EquipBlessNtf{
		HeroIndex: int32(heroIndex),
	}

	randData := make(map[int]int)
	randData[pb.BLESSTYPE_SUCCESS] = blessCfg.Success
	randData[pb.BLESSTYPE_FAIL] = blessCfg.Lose
	randData[pb.BLESSTYPE_INVALID] = blessCfg.Invalid
	randRes := common.RandWeightByMap(randData)

	equipBlessNtf.Res = int32(randRes)
	switch randRes {
	case pb.BLESSTYPE_SUCCESS:
		equip.Lucky += 1
	case pb.BLESSTYPE_FAIL:
		lucky := equip.Lucky - 1
		if lucky < 0 {
			lucky = 0
		}
		equip.Lucky = lucky
	}
	equipBlessNtf.Lucky = int32(equip.Lucky)
	equipBlessNtf.Equip = builder.BuildPbEquipUnit(equip)
	this.GetUserManager().SendMessage(user, equipBlessNtf, true)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}
