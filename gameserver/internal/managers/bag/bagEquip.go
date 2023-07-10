package bag

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

func (this *BagManager) EquipChange(user *objs.User, op *ophelper.OpBagHelperDefault, equipBagPos int, equipTyp int, equip *model.Equip) (*model.Equip, error) {
	var tempEquip *model.Equip
	if equipBagPos != -1 {
		if len(user.Bag) <= equipBagPos || user.Bag[equipBagPos] == nil {
			logger.Error("装备替换，参数错误，装备所在背包中的位置参数异常，目前背包：%v,pos:%v", len(user.Bag), equipBagPos)
			return nil, gamedb.ERRPARAM
		}

		item := user.Bag[equipBagPos]
		itemId := item.ItemId
		tempEquip = user.EquipBag[item.EquipIndex]
		tempEquipT := gamedb.GetEquipEquipCfg(tempEquip.ItemId)
		if tempEquipT == nil || tempEquipT.Type != equipTyp {
			return nil, gamedb.ERREQUIPTYPE
		}
		op.OnGoodsChange(builder.BuildEquipDataChagne(tempEquip.ItemId, -1, 0, equipBagPos, nil), -1)
		BagItemUnitReset(item)
		delete(user.EquipBag, tempEquip.Index)
		//道具记录
		this.GetTlog().ItemFlow(user, itemId, 1, 0, op.GetOpType(), op.OpTypeSecond(), false)
		kyEvent.ItemChange(user, itemId, 1, 0, op.GetOpType(), op.OpTypeSecond(), false)
	}
	if equip.ItemId > 0 {
		if equipBagPos == - 1 {
			equipBagPoses := this.getEmptyPos(user, 1)
			if len(equipBagPoses) == 0 {
				return nil, gamedb.ERRBAGENOUGH
			} else {
				equipBagPos = equipBagPoses[0]
			}
		}
		bagItem(user.Bag[equipBagPos], equip.ItemId, 1)
		user.Bag[equipBagPos].EquipIndex = equip.Index
		user.EquipBag[equip.Index] = equip
		op.OnGoodsChange(builder.BuildEquipDataChagne(equip.ItemId, 1, 1, equipBagPos, equip), 1)
		//道具记录
		this.GetTlog().ItemFlow(user, equip.ItemId, 1, 1, op.GetOpType(), op.OpTypeSecond(), true)
		kyEvent.ItemChange(user, equip.ItemId, 1, 1, op.GetOpType(), op.OpTypeSecond(), true)
	}
	return tempEquip, nil
}

/**
 *  @Description:	装备锁定
 *  @param user
 *  @param position
 */
func (this *BagManager) EquipLock(user *objs.User, position int) error {

	if !this.checkBagPosition(user, constBag.BAG_TYPE_COMMON, position) {
		return gamedb.ERRBAGPOSITION
	}

	item := user.Bag[position]
	if item == nil || item.EquipIndex <= 0 {
		return gamedb.ERRNOTEQUIP
	}

	if user.EquipBag[item.EquipIndex] == nil {
		logger.Error("背包装备数据异常，user:%v,装备：%v", user.Id, item.EquipIndex)
		return gamedb.ERRUNKNOW
	}

	user.EquipBag[item.EquipIndex].IsLock = !user.EquipBag[item.EquipIndex].IsLock
	return nil
}

/**
 *  @Description: 装备回收
 *  @param user
 *  @param positions
 */
func (this *BagManager) EquipRecover(user *objs.User, op *ophelper.OpBagHelperDefault, positions []int32) error {
	addMap := make(map[int]int)
	removePosMap := make(map[int]int)
	userBag, userEquipBag := user.Bag, user.EquipBag
	monthCardPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_EQUIP_RECOVERY)
	for _, v := range positions {
		pos := int(v)
		bagItem := userBag[pos]
		itemId := bagItem.ItemId
		if itemId == 0 {
			continue
		}
		itemBaseCfg := gamedb.GetItemBaseCfg(itemId)
		if itemBaseCfg == nil {
			logger.Error("itemConf error itemId is %v", itemId)
			return gamedb.ERRPARAM
		}
		if len(itemBaseCfg.Recover_items) == 0 {
			return gamedb.ERRNOTRECOVER
		}
		removePosMap[pos] = 0
		for _, award := range itemBaseCfg.Recover_items {
			addMap[award.ItemId] += award.Count * bagItem.Count
		}
	}
	if len(removePosMap) > 0 {
		for pos := range removePosMap {
			bagItem := userBag[pos]
			if bagItem.EquipIndex > 0 {
				op.OnGoodsChange(builder.BuildEquipDataChagne(bagItem.ItemId, -bagItem.Count, 0, pos, userEquipBag[bagItem.EquipIndex]), -bagItem.Count)
				delete(userEquipBag, bagItem.EquipIndex)
			}else{
				op.OnGoodsChange(builder.BuildItemDataChange(bagItem.ItemId, -bagItem.Count, 0, pos), -bagItem.Count)
			}
			this.GetTlog().ItemFlow(user, bagItem.ItemId, bagItem.Count, 0, op.GetOpType(), op.OpTypeSecond(), false)
			kyEvent.ItemChange(user, bagItem.ItemId, bagItem.Count, 0, op.GetOpType(), op.OpTypeSecond(), false)
			BagItemUnitReset(userBag[pos])
		}
	}
	if len(addMap) > 0 {
		addItemInfos := make(gamedb.ItemInfos, 0)
		for itemId, count := range addMap {
			addNum := count
			if monthCardPrivilege != 0 {
				addNum = common.CalcTenThousand(monthCardPrivilege, addNum)
			}
			addItemInfos = append(addItemInfos, &gamedb.ItemInfo{ItemId: itemId, Count: addNum})
		}
		this.AddItems(user, addItemInfos, op)
	}
	this.GetTask().AddTaskProcess(user, pb.CONDITION_RECOVER, 1)
	return nil
}

/**
 *  @Description: 装备销毁
 *  @param user
 *  @param pos 位置  count:数量
 */
func (this *BagManager) EquipDestroy(user *objs.User, op *ophelper.OpBagHelperDefault, pos, count int) error {
	itemId, itemCount := user.Bag[pos].ItemId, user.Bag[pos].Count
	itemBaseCfg := gamedb.GetItemBaseCfg(itemId)
	if itemBaseCfg == nil {
		logger.Error("itemConf error itemId is %v", itemId)
		return gamedb.ERRPARAM
	}
	if count <= 0 {
		return gamedb.ERRPARAM
	}
	if count > user.Bag[pos].Count {
		return gamedb.ERRBAGITEMHAVENOTENOUGHCOUNT
	}
	if len(itemBaseCfg.Destroy_items) == 0 {
		logger.Error("EquipDestroy itemId:%v", itemId)
		return gamedb.ERRNOTDESTROY
	}

	equipIndex := user.Bag[pos].EquipIndex
	lastCount := 0
	if count >= user.Bag[pos].Count {
		//销毁全部
		BagItemUnitReset(user.Bag[pos])
		//装备背包移除
		if equipIndex > 0 {
			delete(user.EquipBag, equipIndex)
		}
		equipIndex = 0
	} else {
		user.Bag[pos].Count -= count
		lastCount = user.Bag[pos].Count
	}
	op1 := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipDestroy1)
	op1.OnGoodsChange(builder.BuildItemDataChange(itemId, -count, lastCount, pos), -count)
	this.GetUserManager().SendItemChangeNtf(user, op1)
	//添加奖励道具
	for _, award := range itemBaseCfg.Destroy_items {
		this.Add(user, op, award.ItemId, award.Count*count)
	}

	//道具记录
	this.GetTlog().ItemFlow(user, itemId, itemCount, lastCount, op.GetOpType(), op.OpTypeSecond(), false)
	kyEvent.ItemChange(user, itemId, itemCount, lastCount, op.GetOpType(), op.OpTypeSecond(), false)
	return nil
}
