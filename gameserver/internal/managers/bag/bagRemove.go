package bag

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"strconv"
)

/**
 *  @Description: 移除ItemInfos道具
 *  @param user
 *  @param op
 *  @param itemInfos
 *  @return error
 */
func (this *BagManager) RemoveItemsInfos(user *objs.User, op *ophelper.OpBagHelperDefault, itemInfos gamedb.ItemInfos) error {
	consumeMap := make(map[int]int)
	has, hasNum := this.HasEnoughItems(user, itemInfos)
	if !has {
		return gamedb.ERRNOTENOUGHGOODS
	}
	for _, info := range itemInfos {
		consumeMap[info.ItemId] += info.Count
	}
	for itemId, count := range consumeMap {
		if count < 1 {
			continue
		}
		_, err := this.IsCheckRemove(user, op, itemId, count, 0, false, false)
		if err != nil {
			return gamedb.ERRNOTENOUGHGOODS
		}

		// 扣除日志
		this.GetTlog().ItemFlow(user, itemId, count, hasNum[itemId]-count, op.GetOpType(), op.OpTypeSecond(), false)
		kyEvent.ItemChange(user, itemId, count, hasNum[itemId]-count, op.GetOpType(), op.OpTypeSecond(), false)
	}
	return nil
}

/**
 *  @Description: 移除道具通用
 *  @param user
 *  @param opHelper
 *  @param itemId
 *  @param count
 *  @return error
 */
func (this *BagManager) Remove(user *objs.User, opHelper *ophelper.OpBagHelperDefault, itemId, count int) error {

	if itemId == pb.ITEMID_RANDOM_STONE || itemId == pb.ITEMID_BACK_CITY {
		return nil
	}

	if count <= 0 {
		logger.Error("玩家：%v,通过：%v 移除道具:%v 数量为0", user.IdName(), opHelper.GetOpType(), itemId)
		return gamedb.ERRITEMZERO
	}

	hasNum, err := this.IsCheckRemove(user, opHelper, itemId, count, 0, true, false)
	if err != nil {
		return err
	}

	// 扣除日志
	this.GetTlog().ItemFlow(user, itemId, count, hasNum-count, opHelper.GetOpType(), opHelper.OpTypeSecond(), false)
	kyEvent.ItemChange(user, itemId, count, hasNum-count, opHelper.GetOpType(), opHelper.OpTypeSecond(), false)
	return nil
}

/**
 *  @Description: 移除道具根据位置
 *  @param user
 *  @param opHelper
 *  @param itemId
 *  @param count
 *  @return error
 */
func (this *BagManager) RemoveByPosition(user *objs.User, opHelper *ophelper.OpBagHelperDefault, itemId, count, pos int) error {

	if itemId == pb.ITEMID_RANDOM_STONE || itemId == pb.ITEMID_BACK_CITY {
		return nil
	}

	if count <= 0 {
		logger.Error("玩家：%v,通过：%v 移除道具:%v 数量为0", user.IdName(), opHelper.GetOpType(), itemId)
		return gamedb.ERRITEMZERO
	}

	hasNum, err := this.IsCheckRemove(user, opHelper, itemId, count, pos, true, true)
	if err != nil {
		return err
	}

	// 扣除日志
	this.GetTlog().ItemFlow(user, itemId, count, hasNum-count, opHelper.GetOpType(), opHelper.OpTypeSecond(), false)
	kyEvent.ItemChange(user, itemId, count, hasNum-count, opHelper.GetOpType(), opHelper.OpTypeSecond(), false)
	return nil
}

func (this *BagManager) IsCheckRemove(user *objs.User, opHelper *ophelper.OpBagHelperDefault, itemId, count, pos int, check, isAuctionRemove bool) (int, error) {
	has, hasNum := true, 0
	if check {
		has, hasNum = this.HasEnough(user, itemId, count)
		if !has {
			logger.Error("玩家：%v,通过：%v 移除道具:%v 数量不足", user.IdName(), opHelper.GetOpType(), itemId)
			return 0, gamedb.ERRNOTENOUGHGOODS
		}
	}

	//非钻石的数量判断
	itemT := gamedb.GetItemBaseCfg(itemId)
	if itemT == nil {
		logger.Error("玩家：%v,通过：%v 添加道具:%v 配置异常", user.IdName(), opHelper.GetOpType(), itemId)
		return 0, gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("item" + strconv.Itoa(itemId))
	}

	var err error
	switch itemT.Type {
	case pb.ITEMTYPE_TOP:
		switch itemId {
		case pb.ITEMID_INGOT:
			this.GetSpendRebates().UpdateSpendRebates(user, opHelper, count)
			this.GetCondition().RecordCondition(user, pb.CONDITION_SPEND_INGOT, []int{count})
		case pb.ITEMID_GOLD:
			this.GetCondition().RecordCondition(user, pb.CONDITION_CONSUME_GOLD, []int{count})
		}

		err = this.topItemRemove(user, opHelper, itemId, count)
	default:
		if isAuctionRemove {
			err = this.normalItemRemoveByPos(user, opHelper, itemId, count, pos)
		} else {
			err = this.normalItemRemove(user, opHelper, itemId, count)
		}
	}
	if err != nil {
		return 0, err
	}
	user.Dirty = true
	return hasNum, err
}

/**
 *  @Description: 普通道具减少
 *  @param user
 *  @param op
 *  @param itemId
 *  @param count
 *  @return error
 */
func (this *BagManager) normalItemRemove(user *objs.User, op *ophelper.OpBagHelperDefault, itemId, count int) error {

	unbindItemId := -1
	bindItemId := -1
	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf.Bind == constBag.Unbind {
		unbindItemId = itemId
	} else {
		bindItemId = itemId
	}
	sameGroupItemId := this.getSameGroupItemId(itemId)
	if sameGroupItemId > 0 {
		sameGroupItemConf := gamedb.GetItemBaseCfg(sameGroupItemId)
		if sameGroupItemConf.Bind == constBag.Unbind {
			unbindItemId = sameGroupItemId
		} else {
			bindItemId = sameGroupItemId
		}
	}

	change := make(map[int]int)
	removeFun := func(removeItemId int) {
		//计算每个位置减少数量
		for i := 0; i < user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum; i++ {

			itemUnitData := user.Bag[i]
			if itemUnitData.ItemId == removeItemId {
				if itemUnitData.Count >= count {
					change[itemUnitData.Position] = count
					count = 0
				} else {
					change[itemUnitData.Position] = itemUnitData.Count
					count -= itemUnitData.Count
				}
				// 已经分配完了
				if count == 0 {
					break
				}
			}
		}
	}
	//优先移除绑定
	if bindItemId > 0 {
		removeFun(bindItemId)
	}
	//移除非绑定
	if unbindItemId > 0 && count > 0 {
		removeFun(unbindItemId)
	}

	if count > 0 {
		return gamedb.ERRNOTENOUGHGOODS
	}

	for pos, value := range change {

		itemUnit := user.Bag[pos]
		itemUnit.Count -= value
		changeItemId := itemUnit.ItemId
		if itemUnit.EquipIndex > 0 {
			delete(user.EquipBag, itemUnit.EquipIndex)
		}
		if itemUnit.Count <= 0 {
			BagItemUnitReset(itemUnit)
		}
		op.OnGoodsChange(builder.BuildItemDataChange(changeItemId, -1*value, itemUnit.Count, pos), -1*value)
	}
	return nil
}

/**
 *  @Description: 普通道具减少根据指定位置
 *  @param user
 *  @param op
 *  @param itemId
 *  @param count
 *  @return error
 */
func (this *BagManager) normalItemRemoveByPos(user *objs.User, op *ophelper.OpBagHelperDefault, itemId, count, pos int) error {

	unbindItemId := -1
	bindItemId := -1
	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf.Bind == constBag.Unbind {
		unbindItemId = itemId
	} else {
		bindItemId = itemId
	}
	sameGroupItemId := this.getSameGroupItemId(itemId)
	if sameGroupItemId > 0 {
		sameGroupItemConf := gamedb.GetItemBaseCfg(sameGroupItemId)
		if sameGroupItemConf.Bind == constBag.Unbind {
			unbindItemId = sameGroupItemId
		} else {
			bindItemId = sameGroupItemId
		}
	}

	change := make(map[int]int)
	removeFun := func(removeItemId int) {
		//计算每个位置减少数量

		itemUnitData := user.Bag[pos]
		if itemUnitData.ItemId == removeItemId {
			if itemUnitData.Count >= count {
				change[itemUnitData.Position] = count
				count = 0
			} else {
				change[itemUnitData.Position] = itemUnitData.Count
				count -= itemUnitData.Count
			}

		}
	}
	//优先移除绑定
	if bindItemId > 0 {
		removeFun(bindItemId)
	}
	//移除非绑定
	if unbindItemId > 0 && count > 0 {
		removeFun(unbindItemId)
	}

	if count > 0 {
		return gamedb.ERRNOTENOUGHGOODS
	}

	for pos, value := range change {
		itemUnit := user.Bag[pos]
		itemUnit.Count -= value
		//changeItemId := itemUnit.ItemId
		if itemUnit.EquipIndex > 0 {
			delete(user.EquipBag, itemUnit.EquipIndex)
		}
		if itemUnit.Count <= 0 {
			BagItemUnitReset(itemUnit)
		}
		equipInfo := user.EquipBag[itemUnit.EquipIndex]

		op.OnGoodsChange(builder.BuildEquipDataChagne(itemId, -value, itemUnit.Count, pos, equipInfo), -value)

		//op.OnGoodsChange(builder.BuildItemDataChange(changeItemId, -1*value, itemUnit.Count, pos), -1*value)
	}
	return nil
}

/**
 *  @Description: 顶级数据减少（元宝 金币） 经验 等级 vip经验 vip等级不能减少
 *  @param user
 *  @param op
 *  @param itemId
 *  @param count
 *  @return error
 */
func (this *BagManager) topItemRemove(user *objs.User, op *ophelper.OpBagHelperDefault, itemId, count int) error {

	now, err := user.AddTopDataByItemId(itemId, -count)
	if err != nil {
		return err
	}
	op.OnGoodsChange(&pb.TopDataChange{Id: int32(itemId), Change: int64(-1 * count), NowNum: int64(now)}, -1*count)

	return nil
}

/**
 *  @Description: 移除指定所有道具
 *  @param user
 *  @param ophelper
 *  @param itemId
 *  @return int
 */
func (this *BagManager) RemoveAllByItemId(user *objs.User, ophelper *ophelper.OpBagHelperDefault, itemId int) int {

	count := 0
	for _, item := range user.Bag {
		if item.ItemId == itemId {
			count += item.Count
			ophelper.OnGoodsChange(builder.BuildItemDataChange(itemId, -1*item.Count, 0, item.Position), -1*item.Count)
			BagItemUnitReset(item)
		}
	}

	// 扣除日志
	this.GetTlog().ItemFlow(user, itemId, count, 0, ophelper.GetOpType(), ophelper.OpTypeSecond(), false)
	kyEvent.ItemChange(user, itemId, count, 0, ophelper.GetOpType(), ophelper.OpTypeSecond(), false)
	return count
}

func (this *BagManager) BagClear(user *objs.User) {

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDebugAddGoods)
	for k, v := range user.Bag {

		if v.ItemId <= 0 {
			continue
		}
		if v.EquipIndex > 0 {
			op.OnGoodsChange(builder.BuildEquipDataChagne(v.ItemId, -v.Count, 0, k, nil), -v.Count)
			delete(user.EquipBag, v.EquipIndex)
		} else {
			op.OnGoodsChange(builder.BuildItemDataChange(v.ItemId, -v.Count, 0, k), -v.Count)
		}
		BagItemUnitReset(v)
	}
	this.GetUserManager().SendItemChangeNtf(user, op)

}
