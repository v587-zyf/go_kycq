package chuanshi

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewChuanShiManager(module managersI.IModule) *ChuanShiManager {
	return &ChuanShiManager{IModule: module}
}

type ChuanShiManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 传世装备穿戴
 *  @param user
 *  @param heroIndex
 *  @param bagPos	背包位置
 *  @param op
 *  @return error
 */
func (this *ChuanShiManager) Wear(user *objs.User, heroIndex, bagPos int, op *ophelper.OpBagHelperDefault, ack *pb.ChuanShiWearAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	bagItemId, err := this.checkBagItem(user, bagPos)
	if err != nil {
		return err
	}
	chuanShiEquipCfg := gamedb.GetChuanShiEquipChuanShiEquipCfg(bagItemId)
	if chuanShiEquipCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, heroIndex, chuanShiEquipCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	equipPos := chuanShiEquipCfg.Type
	heroChuanShi := hero.ChuanShi
	if err = this.GetBag().Remove(user, op, bagItemId, 1); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	if heroChuanShi[equipPos] != 0 {
		this.GetBag().Add(user, op, heroChuanShi[equipPos], 1)
	}
	heroChuanShi[equipPos] = bagItemId
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetCondition().RecordCondition(user, pb.CONDITION_WEAR_CHUAN_SHI_EQUIP, []int{})
	this.GetTask().SendSpecialCheckTaskInfo(user, pb.CONDITION_WEAR_CHUAN_SHI_EQUIP)

	ack.HeroIndex = int32(heroIndex)
	ack.EquipPos = int32(equipPos)
	ack.EquipId = int32(bagItemId)
	return nil
}

/**
 *  @Description: 传世装备卸下
 *  @param user
 *  @param heroIndex
 *  @param equipPos	装备位置
 *  @param op
 *  @return error
 */
func (this *ChuanShiManager) Remove(user *objs.User, heroIndex, equipPos int, op *ophelper.OpBagHelperDefault) error {
	if _, ok := pb.CHUANSHIPOS_MAP[equipPos]; !ok {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if len(hero.ChuanShi) < 1 {
		return gamedb.ERRNOTWEAREQUIP
	}
	equipId := hero.ChuanShi[equipPos]
	if equipId == 0 {
		return gamedb.ERRNOTWEAREQUIP
	}

	this.GetBag().Add(user, op, equipId, 1)
	hero.ChuanShi[equipPos] = 0

	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetCondition().RecordCondition(user, pb.CONDITION_WEAR_CHUAN_SHI_EQUIP, []int{})
	this.GetTask().SendSpecialCheckTaskInfo(user, pb.CONDITION_WEAR_CHUAN_SHI_EQUIP)
	return nil
}

/**
 *  @Description: 传世装备分解(只有背包中道具)
 *  @param user
 *  @param bagPos
 *  @param op
 *  @return error
 */
func (this *ChuanShiManager) DeCompose(user *objs.User, bagPos int, op *ophelper.OpBagHelperDefault) error {
	bagItemId, err := this.checkBagItem(user, bagPos)
	if err != nil {
		return err
	}
	if err = this.GetBag().Remove(user, op, bagItemId, 1); err != nil {
		return err
	}
	chuanShiEquipCfg := gamedb.GetChuanShiEquipChuanShiEquipCfg(bagItemId)
	this.GetBag().AddItems(user, chuanShiEquipCfg.DeComposeItem, op)
	return nil
}

func (this *ChuanShiManager) checkBagItem(user *objs.User, bagPos int) (int, error) {
	bagItemId := this.GetBag().GetItemByPosition(user, bagPos).ItemId
	if bagItemId == 0 {
		return 0, gamedb.ERRGOODS
	}
	itemCfg := gamedb.GetItemBaseCfg(bagItemId)
	if itemCfg == nil || itemCfg.Type != pb.ITEMTYPE_CHUAN_SHI_EQUIP {
		return 0, gamedb.ERRGOODS
	}
	return bagItemId, nil
}
