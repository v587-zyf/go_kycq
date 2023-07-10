package ai

//
//func MonsterDrop(owner base.Actor, monsterId int) ([]*pbserver.ItemUnit, error) {
//
//	monsterConf := gamedb.GetMonsterMonsterCfg(monsterId)
//	//drop, err = GetDropItemList(owner, monsterConf.DropId, monsterConf.Heap)
//	dropItems, err := gamedb.GetDropItems(monsterConf.DropId)
//	if err != nil {
//		return nil, err
//	}
//	if len(dropItems) <= 0 {
//		logger.Error("申请怪物掉落，随机物品为空,:%+v", owner.NickName(), monsterId)
//		return nil, gamedb.ERRUNKNOW
//	}
//
//	drop := make([]*pbserver.ItemUnit, len(dropItems))
//	for k, v := range dropItems {
//		drop[k] = &pbserver.ItemUnit{
//			ItemId:  int32(v.ItemId),
//			ItemNum: int32(v.Count),
//		}
//	}
//
//	specialItem := DropSpacil(owner, monsterConf)
//	if specialItem != nil {
//		drop = append(drop, specialItem)
//	}
//	return drop, nil
//}
//
//func DropSpacil(owner base.Actor, monsterConf *gamedb.MonsterMonsterCfg) *pbserver.ItemUnit {
//	if len(monsterConf.DropSpecial) <= 1 {
//		return nil
//	}
//	userDropInfo := owner.GetFight().GetPlayerDropRedPacketInfo(owner.GetUserId())
//	dropTimes := int(userDropInfo.PickInfos[int32(monsterConf.DropSpecial[0])]) + 1
//
//	dropSpecialItems, err := gamedb.GetDropSpecialItems(monsterConf.Monsterid, dropTimes, int(userDropInfo.PickNum), int(userDropInfo.PickMax))
//	if err != nil {
//		logger.Error("特殊掉落出错 err:%v", err)
//	}
//	if dropSpecialItems == nil {
//		return nil
//	}
//	return &pbserver.ItemUnit{
//		ItemId:  int32(dropSpecialItems.ItemId),
//		ItemNum: int32(dropSpecialItems.Count),
//	}
//}
//
//func GetDropItemList(owner base.Actor, dropId []int, dropNum int) ([]*pbserver.ItemUnit, error) {
//
//	if len(dropId) != 2 {
//		return nil, gamedb.ERRPARAM
//	}
//
//	//获取玩家极品掉落概率
//	bestDropRatio := getUserBestDropRatio(owner, dropId[0])
//	//掉落堆数
//	dropNum = getUserDropNum(owner, dropNum)
//
//	dropConf := gamedb.GetMonsterdropDropCfg(dropId[1])
//	if dropConf == nil {
//		return nil, gamedb.ERRSETTINGNOTFOUND
//	}
//	commonDropWeight := make([]int, len(dropConf.CommonDrop))
//	bestDropWeight := make([]int, len(dropConf.BestDrop))
//	for k, v := range dropConf.CommonDrop {
//		commonDropWeight[k] = v[3]
//	}
//	for k, v := range dropConf.BestDrop {
//		bestDropWeight[k] = v[3]
//	}
//
//	dropItems := make([]*pbserver.ItemUnit, 0)
//
//	for i := 0; i < dropNum; i++ {
//
//		dropConfItems := dropConf.CommonDrop
//		dropWeight := commonDropWeight
//		if len(dropConf.BestDrop) > 0 {
//			if len(dropConfItems) == 0 || common.RandByTenShousand(bestDropRatio) {
//				dropConfItems = dropConf.BestDrop
//				dropWeight = bestDropWeight
//			}
//		}
//
//		dropIndex := common.RandWeightByIntSlice(dropWeight)
//		if dropIndex > -1 && len(dropConfItems) > dropIndex {
//			itemUnit := &pbserver.ItemUnit{
//				ItemId:  int32(dropConfItems[dropIndex][0]),
//				ItemNum: int32(common.RandNum(dropConfItems[dropIndex][1], dropConfItems[dropIndex][2])),
//			}
//			dropItems = append(dropItems, itemUnit)
//		}
//	}
//	return dropItems, nil
//}
//
//func getUserBestDropRatio(owner base.Actor, bestDropRatio int) int {
//
//	return int(float64(bestDropRatio) * 1)
//}
//
//func getUserDropNum(owner base.Actor, dropNum int) int {
//
//	return int(float64(dropNum) * 1)
//}
