package compose

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 合成装备
 *  @param user
 *  @param req
 *  @param op
 *  @return error
 */
func (this *ComposeManager) ComposeEquip(user *objs.User, req *pb.ComposeEquipReq, op *ophelper.OpBagHelperDefault) error {
	composeEquipId := int(req.ComposeEquipSubId)
	composeEquipCfg := gamedb.GetComposeEquipSubComposeEquipSubCfg(composeEquipId)
	if composeEquipId < 1 || composeEquipCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, -1, composeEquipCfg.ComCondition); !check {
		return gamedb.ERRCONDITION
	}
	composeItemInfo := composeEquipCfg.Composeid
	if req.IsLuckyStone {
		if composeEquipCfg.CanLucky == constBag.COMPOSE_LUCKY_NOT {
			return gamedb.ERRCANLUCKY
		}
		//luckConsume := composeEquipCfg.LuckyStone1
		//if req.BigLuckyStone {
		//	luckConsume = composeEquipCfg.LuckyStone2
		//	composeItemInfo = composeEquipCfg.Composeid3
		//} else {
		//	composeItemInfo = composeEquipCfg.Composeid2
		//}
		//if hasEnough, _ := this.GetBag().HasEnough(user, luckConsume.ItemId, luckConsume.Count); luckConsume.ItemId == 0 || !hasEnough {
		//	return gamedb.ERRNOTENOUGHGOODS
		//}
		//removeBagData = append(removeBagData, &gamedb.ItemInfo{ItemId: luckConsume.ItemId, Count: luckConsume.Count})
	}
	cfgNeedMap := make(map[int]int)
	for _, info := range composeEquipCfg.Consume1 {
		cfgNeedMap[info.ItemId] += info.Count
	}
	replaceMap := make(map[int]int)
	removeBagData := make(gamedb.ItemInfos, 0)
	bagPosMap := make(map[int]int)
	reqBagPos := req.GetBagPos()
	if len(reqBagPos) > 0 {
		for _, pos := range req.GetBagPos() {
			bagPos := int(pos)
			bagItem := this.GetBag().GetItemByPosition(user, bagPos)
			itemId, itemCount := bagItem.ItemId, bagItem.Count
			itemCfg := gamedb.GetItemBaseCfg(itemId)
			if _, ok := bagPosMap[bagPos]; ok {
				continue
			}
			if itemCfg == nil || itemCfg.Type != pb.ITEMTYPE_EQUIP {
				logger.Error("客户端发送背包位置物品错误 pos:%v itemId:%v", bagPos, itemId)
				return gamedb.ERRPARAM
			}
			if _, ok := cfgNeedMap[itemId]; ok {
				cfgNeedMap[itemId]--
			}
			if composeEquipCfg := gamedb.GetComposeEquipByComposeId(itemId); composeEquipCfg != nil {
				replaceMap[composeEquipCfg.ReplaceItem.ItemId] += composeEquipCfg.ReplaceItem.Count * itemCount
			}
			bagPosMap[bagPos] = 0
			removeBagData = append(removeBagData, &gamedb.ItemInfo{ItemId: itemId, Count: itemCount})
		}
	}
	userHero := user.Heros
	heroConsumeMap := make(map[int]map[int]int)
	reqEquipPos := req.GetEquipPos()
	if len(reqEquipPos) > 0 {
		for i := 0; i < len(reqEquipPos); i += 2 {
			heroIndex := int(reqEquipPos[i])
			if len(reqEquipPos)-1 < i+1 {
				return gamedb.ERRPARAM
			}
			equipPos := int(reqEquipPos[i+1])
			if heroConsumeMap[heroIndex] == nil {
				heroConsumeMap[heroIndex] = make(map[int]int)
			}
			heroConsumeMap[heroIndex][equipPos] = 0
			heroEquip := userHero[heroIndex].Equips[equipPos]
			if _, ok := cfgNeedMap[heroEquip.ItemId]; ok {
				cfgNeedMap[heroEquip.ItemId]--
			}
			if composeEquipCfg := gamedb.GetComposeEquipByComposeId(heroEquip.ItemId); composeEquipCfg != nil {
				replaceMap[composeEquipCfg.ReplaceItem.ItemId] += composeEquipCfg.ReplaceItem.Count
			}
		}
	}
	reqItems := req.GetItems()
	if len(reqItems) > 0 {
		for i := 0; i < len(reqItems); i += 2 {
			itemId, num := int(reqItems[i]), int(reqItems[i+1])
			if enough, _ := this.GetBag().HasEnough(user, itemId, num); !enough {
				logger.Debug("客户端发送物品数量大于背包物品数量 itemId:%v num:%v", itemId, num)
				return gamedb.ERRNOTENOUGHGOODS
			}
			if needNum, ok := cfgNeedMap[itemId]; ok {
				if needNum < 1 {
					continue
				}
				cfgNeedMap[itemId] -= num
			}
			replaceMap[itemId] += num
			removeBagData = append(removeBagData, &gamedb.ItemInfo{ItemId: itemId, Count: num})
		}
	}
	//把背包碎片取出来
	replaceItemId, replaceNum := composeEquipCfg.ReplaceItem.ItemId, composeEquipCfg.ReplaceItem.Count
	bagNum, _ := this.GetBag().GetItemNum(user, replaceItemId)
	bagReplaceNum := replaceMap[replaceItemId] + bagNum
	logger.Debug("bagReplaceNum:%v 分解Num：%v bagNum:%v replaceNum:%v", bagReplaceNum, replaceMap[replaceItemId], bagNum, replaceNum)
	if bagReplaceNum < replaceNum {
		for itemId, count := range cfgNeedMap {
			if count > 0 {
				logger.Debug("物品不足 itemId:%v count:%v cfgNeedMap:%v", itemId, count, cfgNeedMap)
				return gamedb.ERRNOTENOUGHGOODS
			}
		}
	} else {
		bagReplaceNum -= replaceNum + bagNum
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, removeBagData); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	if len(heroConsumeMap) > 0 {
		for heroIndex, consume := range heroConsumeMap {
			for equipT := range consume {
				userHero[heroIndex].Equips[equipT] = &model.Equip{}
				this.GetUserManager().SendMessage(user, &pb.EquipRemoveAck{
					HeroIndex: int32(heroIndex),
					Pos:       int32(equipT),
				}, true)
			}
		}
		this.GetUserManager().UpdateCombat(user, -1)
	}
	addItems := make(gamedb.ItemInfos, 0)
	addItems = append(addItems, &gamedb.ItemInfo{ItemId: composeItemInfo.ItemId, Count: composeItemInfo.Count})
	if bagReplaceNum > 0 {
		addItems = append(addItems, &gamedb.ItemInfo{ItemId: replaceItemId, Count: bagReplaceNum})
	}
	kyEvent.Compose(user, composeItemInfo.ItemId, composeItemInfo.Count, removeBagData)

	//composeMap := make(map[int]int)
	//randMap := make(map[int]int)
	//for _, ints := range composeItemInfo {
	//	composeMap[ints[0]] = ints[1]
	//	randMap[ints[0]] = ints[2]
	//}
	//composeId := common.RandWeightByMap(randMap)
	this.GetBag().AddItems(user, addItems, op)
	user.Dirty = true
	return nil
}
