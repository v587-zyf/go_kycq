package fit

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"math"
)

const (
	FITHOLYEQUIP_COMPOSE_NOTCHECK_GRADE = 1
)

/**
 *  @Description: 合体圣装合成、升级
 *  @param user
 *  @param equipId	装备id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) HolyEquipCompose(user *objs.User, equipId, equipPos int, op *ophelper.OpBagHelperDefault, ack *pb.FitHolyEquipComposeAck) error {
	fitHolyEquipCfg := gamedb.GetFitHolyEquipFitHolyEquipCfg(equipId)
	if fitHolyEquipCfg == nil {
		return gamedb.ERRPARAM
	}
	switch equipPos {
	case pb.EQUIPPOS_NINE:
		if pb.EQUIPPOS_FIVE != fitHolyEquipCfg.Type {
			return gamedb.ERRPARAM
		}
	case pb.EQUIPPOS_TEN:
		if pb.EQUIPPOS_SIX != fitHolyEquipCfg.Type {
			return gamedb.ERRPARAM
		}
	default:
		if equipPos != fitHolyEquipCfg.Type {
			return gamedb.ERRPARAM
		}
	}
	if !this.GetCondition().CheckMulti(user, -1, fitHolyEquipCfg.Condition) {
		return gamedb.ERRCONDITION
	}
	userFitHolyEquip := user.FitHolyEquip

	suitT := fitHolyEquipCfg.SuitType
	if userFitHolyEquip.Equips[suitT] == nil {
		userFitHolyEquip.Equips[suitT] = make(model.IntKv)
	}
	fitHolyEquips := userFitHolyEquip.Equips[suitT]
	if fitHolyEquipCfg.Grade != FITHOLYEQUIP_COMPOSE_NOTCHECK_GRADE {
		tempPos := equipPos
		switch equipPos {
		case pb.EQUIPPOS_NINE:
			tempPos = pb.EQUIPPOS_FIVE
		case pb.EQUIPPOS_TEN:
			tempPos = pb.EQUIPPOS_SIX
		}
		if gamedb.GetFitHolyEquipByPosAndGrade(suitT, tempPos, fitHolyEquipCfg.Grade-1).Id != fitHolyEquips[equipPos] {
			return gamedb.ERRPARAM
		}
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, fitHolyEquipCfg.ComposeItem); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	fitHolyEquips[equipPos] = equipId
	kyEvent.FitHolyEquipLvUp(user, suitT, equipId, fitHolyEquipCfg.Grade)

	ack.SuitType = int32(suitT)
	ack.FitHolyEquip = builder.BuildFitHolyEquipUnit(fitHolyEquips)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 合体圣装分解
 *  @param user
 *  @param bagPos	背包位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) HolyEquipDeCompose(user *objs.User, bagPos int, op *ophelper.OpBagHelperDefault, ack *pb.FitHolyEquipDeComposeAck) error {
	itemId, fitHolyEquipCfg, err := this.CheckBagItem(user, bagPos)
	if err != nil {
		return err
	}
	if err = this.GetBag().Remove(user, op, itemId, 1); err != nil {
		return err
	}
	this.GetBag().AddItems(user, fitHolyEquipCfg.DeComposeItem, op)

	ack.Goods = op.ToChangeItems()
	return nil
}

/**
 *  @Description: 合体圣装穿戴
 *  @param user
 *  @param bagPos	背包位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) HolyEquipWear(user *objs.User, bagPos, equipPos int, op *ophelper.OpBagHelperDefault, ack *pb.FitHolyEquipWearAck) error {
	itemId, fitHolyEquipCfg, err := this.CheckBagItem(user, bagPos)
	if err != nil {
		return err
	}
	if !this.GetCondition().CheckMulti(user, -1, fitHolyEquipCfg.Condition) {
		return gamedb.ERRCONDITION
	}
	if err = this.GetBag().RemoveByPosition(user, op, itemId, 1, bagPos); err != nil {
		return err
	}

	suitT, equipType := fitHolyEquipCfg.SuitType, fitHolyEquipCfg.Type
	switch equipPos {
	case pb.EQUIPPOS_NINE:
		if pb.EQUIPPOS_FIVE != equipType {
			return gamedb.ERRPARAM
		}
	case pb.EQUIPPOS_TEN:
		if pb.EQUIPPOS_SIX != equipType {
			return gamedb.ERRPARAM
		}
	default:
		if equipPos != equipType {
			return gamedb.ERRPARAM
		}
	}
	userFitHolyEquip := user.FitHolyEquip
	userFitHolyEquips, ok := userFitHolyEquip.Equips[suitT]
	if !ok {
		userFitHolyEquip.Equips[suitT] = make(model.IntKv)
		userFitHolyEquips = userFitHolyEquip.Equips[suitT]
	}
	if userFitHolyEquips[equipPos] != 0 {
		this.GetBag().Add(user, op, userFitHolyEquips[equipPos], 1)
	}
	userFitHolyEquips[equipPos] = itemId

	ack.SuitType = int32(suitT)
	ack.FitHolyEquip = builder.BuildFitHolyEquipUnit(userFitHolyEquips)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 合体圣装卸下
 *  @param user
 *  @param pos	装备位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) HolyEquipRemove(user *objs.User, pos, suitType int, op *ophelper.OpBagHelperDefault, ack *pb.FitHolyEquipRemoveAck) error {
	if _, ok := pb.FITHOLYEQUIPPOS_MAP[pos]; !ok {
		return gamedb.ERRPARAM
	}
	userFitHolyEquips, ok := user.FitHolyEquip.Equips[suitType]
	if !ok {
		return gamedb.ERRPARAM
	}
	if userFitHolyEquips[pos] != 0 {
		this.GetBag().Add(user, op, userFitHolyEquips[pos], 1)
	}
	userFitHolyEquips[pos] = 0

	ack.SuitType = int32(suitType)
	ack.FitHolyEquip = builder.BuildFitHolyEquipUnit(userFitHolyEquips)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 合体圣装套装技能更换
 *  @param user
 *  @param suitId	套装id
 *  @return error
 */
func (this *FitManager) HolyEquipSuitSkillChange(user *objs.User, suitId int) error {
	fitHolyEquipSuitCfg := gamedb.GetFitHolyEquipSuitFitHolyEquipSuitCfg(suitId)
	if fitHolyEquipSuitCfg == nil {
		return gamedb.ERRPARAM
	}
	userFitHolyEquip := user.FitHolyEquip
	minGrade, isSuit := math.MaxInt32, true
	suitT, suitGrade := suitId/constConstant.COMPUTE_TEN_THOUSAND, suitId%constConstant.COMPUTE_TEN_THOUSAND
	for _, id := range userFitHolyEquip.Equips[suitT] {
		fitHolyCfg := gamedb.GetFitHolyEquipFitHolyEquipCfg(id)
		if id <= 0 || fitHolyCfg == nil {
			isSuit = false
			continue
		}
		if minGrade > fitHolyCfg.Grade {
			minGrade = fitHolyCfg.Grade
		}
	}
	if !isSuit || minGrade < suitGrade {
		return gamedb.ERRSUITNOTACTIVE
	}
	userFitHolyEquip.SuitId = suitId
	kyEvent.FitHolyEquipSuitSkillChange(user, suitT, suitGrade)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

func (this *FitManager) CheckBagItem(user *objs.User, bagPos int) (int, *gamedb.FitHolyEquipFitHolyEquipCfg, error) {
	bagItem := this.GetBag().GetItemByPosition(user, bagPos)
	if bagItem.ItemId == 0 {
		return 0, nil, gamedb.ERRPARAM
	}
	itemId := bagItem.ItemId
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	if itemCfg == nil || itemCfg.Type != pb.ITEMTYPE_FIT_HOLY_EQUIP {
		return 0, nil, gamedb.ERRGOODS
	}
	fitHolyEquipCfg := gamedb.GetFitHolyEquipFitHolyEquipCfg(itemId)
	return itemId, fitHolyEquipCfg, nil
}
