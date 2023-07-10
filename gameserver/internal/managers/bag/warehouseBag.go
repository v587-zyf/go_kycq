package bag

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

func (this *BagManager) WareHouseOnline(user *objs.User) {

	if user.WarehouseBagInfo == nil {
		user.WarehouseBagInfo = make(map[int]*model.BagInfoUnit)
	}

	//初始化玩家背包格数
	if user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON] == nil {
		user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON] = &model.BagInfoUnit{
			MaxNum:   gamedb.GetWareHouseBagInitNum(),
			SpaceAdd: make(model.IntKv),
			BuyNum:   0,
		}
	}
	logger.Debug("gamedb.GetWareHouseBagInitNum():%v", gamedb.GetWareHouseBagInitNum())
	emptyPos := make(map[int]bool)
	for i := 0; i < user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum; i++ {
		emptyPos[i] = true
	}
	//记录玩家已背包格子数据
	bagData := make(model.Bag, user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum)
	for _, v := range user.WarehouseBag {
		if v.Position >= len(bagData) {
			logger.Info("WareHouseOnline v.position:%v  MaxNum:%v", v.Position, user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum)
			continue
		}
		bagData[v.Position] = v
		delete(emptyPos, v.Position)
	}
	user.WarehouseBag = bagData
	//初始化玩家背包格子数据
	for k := range emptyPos {
		user.WarehouseBag[k] = &model.Item{Position: k}
		bagItem(user.WarehouseBag[k], 0, 0)
	}
}

func (this *BagManager) WareHouseBagSpaceAdd(user *objs.User, op *ophelper.OpBagHelperDefault) error {

	buyNum := user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].BuyNum
	if buyNum >= gamedb.GetBagAddCfgLen(constBag.WAREHOUSE_BAG_ADD_TYPE_ITEM) {
		return gamedb.ERRCANNOTBUYSPACE
	}
	costItemId, costCount, addNum := gamedb.GetWarehouseBagSpaceAddCost(buyNum)
	hasEnough, _ := this.HasEnough(user, costItemId, costCount)
	if !hasEnough {
		return gamedb.ERRNOTENOUGHGOODS
	}
	err := this.Remove(user, op, costItemId, costCount)
	if err != nil {
		return err
	}
	this.warehouseBagSpaceAdd(user, addNum)
	user.Dirty = true
	return nil
}

func (this *BagManager) warehouseBagSpaceAdd(user *objs.User, addNum int) {
	nowNum := user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum
	maxNum := nowNum + addNum
	user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum = maxNum
	for ; nowNum < maxNum; nowNum++ {
		user.WarehouseBag = append(user.WarehouseBag, &model.Item{Position: nowNum})
	}
	user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].BuyNum += 1
}

//
//  WareHouseBagAdd
//  @Description:  移动背包物品到仓库
//
func (this *BagManager) WareHouseBagAdd(user *objs.User, op *ophelper.OpBagHelperDefault, positions []int32) error {

	err := this.BagMoveChangeBeforeCheck(positions, user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum)
	if err != nil {
		logger.Error("WareHouseBagAdd userId:%v  pos:%v", user.Id, positions)
		return err
	}

	isHaveEnoughSpace := this.checkIsHaveEnoughSpace(user.WarehouseBag, len(positions))
	if !isHaveEnoughSpace {
		return gamedb.ERRWAREHOUSEHAVENOENOUGHSPACE
	}

	//校验是否可移动
	for _, v := range positions {
		itemId := user.Bag[int(v)].ItemId
		itemBaseCfg := gamedb.GetItemBaseCfg(itemId)
		if itemBaseCfg == nil {
			logger.Error("itemConf error itemId is %v", itemId)
			return gamedb.ERRPARAM
		}
	}
	data := make(map[int]*model.MoveInfo)
	//移动背包数据到仓库
	for _, v := range positions {
		pos := int(v)
		itemId := user.Bag[pos].ItemId
		count := user.Bag[pos].Count
		equipIndex := user.Bag[pos].EquipIndex
		logger.Debug(" 移入仓库   userId:%v  ItemId:%v  equipIndex:%v", user.Id, itemId, equipIndex)
		//背包移除道具
		BagItemUnitReset(user.Bag[pos])
		//仓库增加道具
		this.AddItemsToBag(user.WarehouseBag, user.Id, itemId, count, equipIndex, data)
		if equipIndex > 0 {
			op.OnGoodsChange(builder.BuildEquipDataChagne(itemId, -count, 0, pos, user.EquipBag[equipIndex]), -count)
		} else {
			op.OnGoodsChange(builder.BuildItemDataChange(itemId, -count, 0, pos), -count)
		}

		//道具记录
		this.GetTlog().ItemFlow(user, itemId, count, 0, op.GetOpType(), op.OpTypeSecond(), false)
		kyEvent.ItemChange(user, itemId, count, 0, op.GetOpType(), op.OpTypeSecond(), false)
	}

	return nil
}

//
//  WareHouseMoveToBag
//  @Description: 移出仓库物品到背包
//
func (this *BagManager) WareHouseMoveToBag(user *objs.User, op *ophelper.OpBagHelperDefault, positions []int32) error {

	err := this.BagMoveChangeBeforeCheck(positions, user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum)
	if err != nil {
		logger.Error("WareHouseMoveToBag userId:%v  pos:%v", user.Id, positions)
		return err
	}

	isHaveEnoughSpace := this.checkIsHaveEnoughSpace(user.Bag, len(positions))
	if !isHaveEnoughSpace {
		return gamedb.ERRWAREHOUSEHAVENOENOUGHSPACE
	}

	//校验是否可移动
	for _, v := range positions {
		if user.WarehouseBag[int(v)] == nil {
			return gamedb.ERRPARAM
		}
		itemId := user.WarehouseBag[int(v)].ItemId
		itemBaseCfg := gamedb.GetItemBaseCfg(itemId)
		if itemBaseCfg == nil {
			logger.Error("itemConf error itemId is %v", itemId)
			return gamedb.ERRPARAM
		}
	}

	//移动厂库数据到背包
	data := make(map[int]*model.MoveInfo)
	hasNumMap := make(map[int]int)
	for _, v := range positions {
		pos := int(v)
		itemId := user.WarehouseBag[pos].ItemId
		count := user.WarehouseBag[pos].Count
		equipIndex := user.WarehouseBag[pos].EquipIndex
		logger.Debug(" 移出仓库   userId:%v  ItemId:%v  equipIndex:%v", user.Id, itemId, equipIndex)
		//仓库移除道具
		BagItemUnitReset(user.WarehouseBag[pos])
		//背包增加道具
		itemCfg := gamedb.GetItemBaseCfg(itemId)
		bagPos, _, dataInfos := this.AddItemsToBag(user.Bag, user.Id, itemId, count, equipIndex, data)
		
		if itemCfg.Type != pb.ITEMTYPE_EQUIP && itemCfg.CountLimit >= 1 {
			data = dataInfos
		} else {
			op.OnGoodsChange(builder.BuildEquipDataChagne(itemId, count, count, bagPos, user.EquipBag[equipIndex]), count)
		}

		//道具记录
		hasNum, ok := hasNumMap[itemId]
		if !ok {
			hasNum, _ = this.GetItemNum(user, itemId)
		} else {
			hasNum += count
		}
		this.GetTlog().ItemFlow(user, itemId, count, hasNum, op.GetOpType(), op.OpTypeSecond(), true)
		kyEvent.ItemChange(user, itemId, count, hasNum, op.GetOpType(), op.OpTypeSecond(), true)
	}

	for pos, v := range data {
		logger.Debug(" pos:%v  非装备道具移动信息:%v", pos, v)
		changeCount := v.Count
		if v.IsNewPos && v.AllCount == v.CountLimit {
			changeCount = v.AllCount
		}
		op.OnGoodsChange(builder.BuildItemDataChange(v.ItemId, changeCount, v.AllCount, pos), changeCount)
	}

	return nil
}

//
//  checkIsHaveEnoughSpace
//  @Description: 检查仓库空间是否足够
//
func (this *BagManager) checkIsHaveEnoughSpace(bagInfo model.Bag, lenPositions int) bool {
	//true :仓库空间足够  false:仓库空间不够
	canAddNum := 0
	for _, v := range bagInfo {
		if v.ItemId == 0 {
			canAddNum++
		}
	}

	if lenPositions > canAddNum {
		return false
	}
	return true
}

//增加道具到仓库
func (this *BagManager) AddItemsToWareBag(user *objs.User, itemId, count, equipIndex int) {
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		logger.Error("配置错误 itemId:%v", itemId)
		return
	}

	dataNum := make(map[int]int)
	if cfg.Type == pb.ITEMTYPE_EQUIP {
		for index, v := range user.WarehouseBag {
			if v.ItemId == 0 {
				v.ItemId = itemId
				v.Count = count
				v.EquipIndex = equipIndex
				v.Position = index
				break
			}
		}
		return
	} else {
		for _, v := range user.WarehouseBag {
			dataNum[v.ItemId] = v.Count
		}

		for index, v := range user.WarehouseBag {
			allNum := dataNum[itemId]
			otherAdd := 0
			if allNum > 0 {
				if allNum+count <= cfg.CountLimit {
					if v.ItemId == itemId {
						v.Count += count
						break
					}
				} else {
					otherAdd = allNum + count - cfg.CountLimit
					if v.ItemId == itemId {
						v.Count = cfg.CountLimit
					}
					for index, v := range user.WarehouseBag {
						if v.ItemId == 0 {
							v.ItemId = itemId
							v.Count = otherAdd
							v.EquipIndex = equipIndex
							v.Position = index
							break
						}
					}
				}
			} else {
				if v.ItemId == 0 {
					v.ItemId = itemId
					v.Count = count
					v.EquipIndex = equipIndex
					v.Position = index
					break
				}
			}
		}
	}
}

//移动道具到背包
func (this *BagManager) AddItemsToBag(bagInfo model.Bag, userId, itemId, count, equipIndex int, data map[int]*model.MoveInfo) (int, int, map[int]*model.MoveInfo) {
	pos := 0
	beforeCount := 0
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	if itemCfg.Type != pb.ITEMTYPE_EQUIP && itemCfg.CountLimit >= 1 {
		for index, v := range bagInfo {
			if v.Count == itemCfg.CountLimit {
				continue
			}
			if v.ItemId == itemId && v.Count < itemCfg.CountLimit {
				beforeCount = v.Count
				if count+v.Count <= itemCfg.CountLimit {
					v.ItemId = itemId
					v.Count += count
					v.EquipIndex = equipIndex
					v.Position = index
					data = this.setDataInfo(data, v.Position, itemId, count, equipIndex, beforeCount, count+beforeCount, itemCfg.CountLimit, false)
					logger.Debug("AddItemsToBag userId:%v  itemId:%v  count:%v  bagCount:%v  index:%v pos:%v", userId, itemId, count, v.Count, index, pos)
					return pos, beforeCount, data
				} else {
					beforeNum := v.Count
					addCount := count + v.Count - itemCfg.CountLimit
					v.ItemId = itemId
					v.Count = itemCfg.CountLimit
					v.EquipIndex = equipIndex
					v.Position = index
					pos = v.Position
					if data[pos] == nil {
						data[pos] = &model.MoveInfo{}
					}
					data[pos].ItemId = itemId
					data[pos].Count = itemCfg.CountLimit - beforeNum
					data[pos].EquipIndex = equipIndex
					data[pos].BeforeCount = beforeCount
					data[pos].AllCount = itemCfg.CountLimit
					data[pos].CountLimit = itemCfg.CountLimit

					logger.Debug("AddItemsToBag 超格子单个数量上限 userId:%v  itemId:%v  count:%v  index:%v pos:%v,addCount:%v", userId, itemId, count, index, pos, addCount)

					for _, v1 := range bagInfo {
						if v1.Count == itemCfg.CountLimit {
							continue
						}
						if v1.ItemId == itemId && v1.Count < itemCfg.CountLimit {
							pos = v1.Position
							if addCount+v1.Count >= itemCfg.CountLimit {
								data = this.setDataInfo(data, v1.Position, itemId, itemCfg.CountLimit-v1.Count, equipIndex, beforeCount, itemCfg.CountLimit, itemCfg.CountLimit, false)
								addCount = addCount + v1.Count - itemCfg.CountLimit
								v1.Count = itemCfg.CountLimit
							} else {
								data = this.setDataInfo(data, v1.Position, itemId, addCount, equipIndex, beforeCount, addCount+v1.Count, itemCfg.CountLimit, false)
								v1.Count += addCount
								addCount = 0
							}
						}
					}

					if addCount > 0 {
						for index, v := range bagInfo {
							if v.ItemId == 0 {
								v.ItemId = itemId
								v.Count = addCount
								v.EquipIndex = equipIndex
								v.Position = index
								data = this.setDataInfo(data, v.Position, itemId, addCount, equipIndex, 0, addCount, itemCfg.CountLimit, true)
								logger.Debug("AddItemsToBag 非装备类型 叠加数量超上限 新开一个格子 userId:%v  itemId:%v  count:%v  index:%v pos:%v", userId, itemId, count, index, pos)
								break
							}
						}
						return pos, beforeCount, data
					}
				}
			}
		}
	}
	for index, v := range bagInfo {
		if v.ItemId == 0 {
			v.ItemId = itemId
			v.Count = count
			v.EquipIndex = equipIndex
			v.Position = index
			pos = v.Position
			if itemCfg.Type != pb.ITEMTYPE_EQUIP && itemCfg.CountLimit >= 1 {
				data = this.setDataInfo(data, v.Position, itemId, count, equipIndex, 0, count, itemCfg.CountLimit, true)
			}
			logger.Debug("AddItemsToBag userId:%v  itemId:%v  count:%v  index:%v pos:%v", userId, itemId, count, index, pos)
			break
		}
	}
	return pos, beforeCount, data
}

func (this *BagManager) setDataInfo(data map[int]*model.MoveInfo, pos, itemId, count, equipIndex, beforeCount, allCount, countLimit int, isNewPos bool) map[int]*model.MoveInfo {

	if data[pos] == nil {
		data[pos] = &model.MoveInfo{}
	}
	data[pos].ItemId = itemId
	data[pos].Count = count
	data[pos].EquipIndex = equipIndex
	data[pos].BeforeCount = beforeCount
	data[pos].AllCount = allCount
	data[pos].CountLimit = countLimit
	data[pos].IsNewPos = isNewPos
	return data
}

//物品移入  移出 仓库检查
func (this *BagManager) BagMoveChangeBeforeCheck(pos []int32, bagCount int) error {
	data := make(map[int32]bool)
	for _, p := range pos {
		if p > int32(bagCount-1) || p < 0 {
			return gamedb.ERRPARAM
		}
		if data[p] {
			return gamedb.ERRPARAM
		}
		data[p] = true
	}
	return nil
}
