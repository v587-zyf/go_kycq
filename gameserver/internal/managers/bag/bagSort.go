package bag

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"sort"
	"time"
)

/*
比角色身上更好的装备 >道具>同级别或更差的装备>非本职业的装备
		第一规则	装备好坏的对比规则
			1.根据装备id，在equip表中对比同部位装备的阶数，阶数大的装备更好
			2.阶数相同时，对比品质，品质高的装备更好
			3.品质相同时，对比星级，星级高的装备更好

		第二规则	相同优先级时的排序规则
			1.按item表的itemValues的大小排序，小的在前，大的在后
			2.value相同的，按item表quality排序，品质好的在前，品质差的在后
			3.quality相同的，按item表的id排序，小的在前，大的在后

			是否是本职业装备的判断规则
			1.根据equip表的condition字段判断

	//=========新规则 背包里装备战力 比 穿戴的装备战力高的 优先排在前面 (未激活的hero 的装备不参与此次排序排序)  代替之前根据品质排的规则！！！ ==========
*/
func (this *BagManager) BagSort(user *objs.User, isWarehouse bool) error {
	bag := user.Bag
	bagTypes := constBag.BAG_TYPE_COMMON
	if isWarehouse {
		bag = user.WarehouseBag
		bagTypes = constBag.WAREHOUSE_BAG_TYPE_COMMON
	}
	if bag == nil {
		return nil
	}
	if user.SortTimes == nil {
		user.SortTimes = make(model.IntKv)
	}
	if int64(user.SortTimes[bagTypes]+gamedb.GetConf().BagInterval) > time.Now().Unix() {
		return gamedb.ERRBAGSORT
	}
	user.SortTimes[bagTypes] = int(time.Now().Unix())

	allBagSuperposition := make(map[int]*model.ItemMark, 0)
	randNum := 19960806
	//======== 排序前先把背包非装备物品自动堆叠
	for _, item := range bag {
		if item == nil {
			continue
		}
		if item.ItemId == 0 {
			continue
		}
		itemConf := gamedb.GetItemBaseCfg(item.ItemId)
		if itemConf != nil {
			if itemConf.CountLimit == 0 {
				itemConf.CountLimit = 999
			}
			//装备不叠加
			if itemConf.Type == pb.ITEMTYPE_EQUIP {
				allBagSuperposition[item.ItemId+randNum] = &model.ItemMark{ItemId: item.ItemId, Count: item.Count, EquipIndex: item.EquipIndex, CountLimit: itemConf.CountLimit, ItemType: itemConf.Type}
				randNum++
			} else {
				if allBagSuperposition[item.ItemId] == nil {
					allBagSuperposition[item.ItemId] = &model.ItemMark{}
				}
				count := allBagSuperposition[item.ItemId].Count
				allBagSuperposition[item.ItemId] = &model.ItemMark{ItemId: item.ItemId, Count: item.Count + count, EquipIndex: item.EquipIndex, CountLimit: itemConf.CountLimit, ItemType: itemConf.Type}
			}
		}
	}

	for _, data := range allBagSuperposition {
		logger.Debug("allBagSuperposition itemId:%v  count:%v  countLimit:%v", data.ItemId, data.Count, data.CountLimit)
	}

	bag = this.ResetBag(user, isWarehouse, nil, nil, nil, allBagSuperposition)

	//=========新规则 背包里装备战力 比 穿戴的装备战力高的 优先排在前面 (未激活的hero 的装备不参与此次排序排序)==========
	allOverWearCombatEquip := make(map[int]model.ItemMarkSlice, 0) //背包里超过身上穿戴装备的战力排序
	minWearEquipCombat := make(map[int]map[int]int)                //身上穿戴的最小战力的装备
	activityJobMark := make(map[int]bool)
	for _, heroInfo := range user.Heros {
		if minWearEquipCombat[heroInfo.Job] == nil {
			minWearEquipCombat[heroInfo.Job] = make(map[int]int)
		}
		activityJobMark[heroInfo.Job] = true
		for _, v := range heroInfo.Equips {
			if v != nil {
				if v.ItemId > 0 {
					cfg := gamedb.GetEquipEquipCfg(v.ItemId)
					if cfg != nil {
						combat := builder.CalcCombat(heroInfo.Job, v.ItemId, -1, cfg.Properties, v.RandProps)
						logger.Debug("背包整理(一)  玩家身上穿戴的装备 isWarehouse:%v   userId:%v   itemType:%v  combat:%v  itemId:%v  heroJob:%v  cfg.Properties:%v v.RandProps:%v", isWarehouse, user.Id, cfg.Type, combat, v.ItemId, heroInfo.Job, cfg.Properties, v.RandProps)
						if minWearEquipCombat[heroInfo.Job][cfg.Type] <= 0 {
							minWearEquipCombat[heroInfo.Job][cfg.Type] = combat
						} else {
							if combat < minWearEquipCombat[heroInfo.Job][cfg.Type] {
								minWearEquipCombat[heroInfo.Job][cfg.Type] = combat
							}
						}
					}
				}
			}
		}
	}
	logger.Debug("BagSort 背包整理(二)   userId:%v isWarehouse:%v  各个部位身上穿戴装备的最小战力 minWearEquipCombat:%v", user.Id, isWarehouse, minWearEquipCombat)
	for index, item := range bag {
		if item == nil {
			continue
		}
		if item.ItemId == 0 {
			continue
		}
		itemConf := gamedb.GetItemBaseCfg(item.ItemId)
		if itemConf != nil {
			if itemConf.Type == pb.ITEMTYPE_EQUIP { //是装备类型道具
				//判断是否是当前激活的职业能佩戴的道具
				state, cfg := this.GetCondition().CheckItemIsCanBeUse(user, item.ItemId)
				if state {
					randomProp := make([]*model.EquipRandProp, 0)
					if user.EquipBag[item.EquipIndex] != nil {
						randomProp = user.EquipBag[item.EquipIndex].RandProps
					}
					calcCombat := builder.CalcCombat(cfg.Condition[pb.CONDITION_JOB], item.ItemId, index, cfg.Properties, randomProp)
					if !activityJobMark[cfg.Condition[pb.CONDITION_JOB]] {
						continue
					}
					if calcCombat > minWearEquipCombat[cfg.Condition[pb.CONDITION_JOB]][cfg.Type] {
						//=========新规则 背包里装备战力 比 穿戴的装备战力高的 优先排在前面 (未激活的hero 的装备不参与此次排序排序)==========
						if allOverWearCombatEquip[cfg.Type] == nil {
							allOverWearCombatEquip[cfg.Type] = make(model.ItemMarkSlice, 0)
						}
						allOverWearCombatEquip[cfg.Type] = append(allOverWearCombatEquip[cfg.Type], &model.ItemMark{ItemId: item.ItemId, Count: item.Count, Index: index, Class: cfg.Class, Quality: cfg.Quality, Star: cfg.Star, EquipIndex: item.EquipIndex, Combat: calcCombat})
					}
				}
			}
		} else {
			logger.Error("gamedb.GetItemBaseCfg(item.ItemId)  cfg nil:%v", item.ItemId)
		}
	}

	allBagAndWearEquip := make(model.ItemMarkSlice3, 0) //所有背包里比自己身上穿戴好的的装备
	for equipType, itemInfo := range allOverWearCombatEquip {
		for _, v := range itemInfo {
			if v.Index == -1 {
				// index == -1 说明是玩家身上装备的道具
				break
			}
			itemConf := gamedb.GetItemBaseCfg(v.ItemId)
			if itemConf != nil {
				allBagAndWearEquip = append(allBagAndWearEquip, &model.ItemMark{ItemId: v.ItemId, Count: v.Count, Index: v.Index, Class: v.Class, Quality: v.Quality, Star: v.Star, ItemValue: itemConf.ItemValues, ItemQuality: itemConf.Quality, ItemCfgId: itemConf.Id, EquipIndex: v.EquipIndex, Combat: v.Combat, EquipType: equipType})
			}
		}
	}
	sort.Sort(allBagAndWearEquip)
	//-------- 日志打印
	for _, v := range allBagAndWearEquip {
		logger.Debug("背包整理(三) isWarehouse:%v  所有比身上穿戴装备战力高的  背包里的装备的 排序    itemId:%v  index:%v  class:%v quality:%v star:%v   itemValue:%v  itemQuality:%v itemCfgId:%v  equipType:%v  combat:%v", isWarehouse, v.ItemId, v.Index, v.Class, v.Quality, v.Star, v.ItemValue, v.ItemQuality, v.ItemCfgId, v.EquipType, v.Combat)
	}
	//-------- 日志打印

	//装备排完序了  开始排背包里剩下的  (非装备东西和比自己身上佩戴差的装备) 放在一起 再按照第二规则排序
	haveSortBagIndex := make(map[int]bool) //已经排好序的背包装备 key bag的index下标
	for _, v := range allBagAndWearEquip {
		if v.Index == -1 {
			continue
		}
		haveSortBagIndex[v.Index] = true
	}

	//背包里非装备的物品 and 比自己身上差的装备道具
	allNotEquipAndBadEquipItems := make(model.ItemMarkSlice2, 0)

	//所有比身上装备差的装备
	allLowEquip := make(model.ItemMarkSlice2, 0)

	for index, v := range bag {
		if haveSortBagIndex[index] {
			//已经排过序
			continue
		}
		if v == nil {
			continue
		}
		itemConf := gamedb.GetItemBaseCfg(v.ItemId)
		if itemConf != nil {
			if itemConf.Type == pb.ITEMTYPE_EQUIP {
				allLowEquip = append(allLowEquip, &model.ItemMark{ItemId: v.ItemId, Count: v.Count, Index: index, ItemValue: itemConf.ItemValues, ItemQuality: itemConf.Quality, ItemCfgId: itemConf.Id, EquipIndex: v.EquipIndex})
			} else {
				allNotEquipAndBadEquipItems = append(allNotEquipAndBadEquipItems, &model.ItemMark{ItemId: v.ItemId, Count: v.Count, Index: index, ItemValue: itemConf.ItemValues, ItemQuality: itemConf.Quality, ItemCfgId: itemConf.Id, EquipIndex: v.EquipIndex})
			}
		}
	}
	sort.Sort(allNotEquipAndBadEquipItems)
	//-------- 日志打印
	for _, v := range allNotEquipAndBadEquipItems {
		logger.Debug("背包整理(四) isWarehouse:%v 背包里比自己当前身上佩戴差的装备和非装备道具放到一个池子里排完序后的顺序    itemId:%v  class:%v quality:%v star:%v   itemValue:%v  itemQuality:%v itemCfgId:%v", isWarehouse, v.ItemId, v.Class, v.Quality, v.Star, v.ItemValue, v.ItemQuality, v.ItemCfgId)
	}
	//-------- 日志打印

	//-------- 日志打印
	for _, v := range allLowEquip {
		logger.Debug("背包整理(四) isWarehouse:%v 背包里比自己当前身上佩戴差的装备和未激活战士装备道具放到最后排序    itemId:%v  class:%v quality:%v star:%v   itemValue:%v  itemQuality:%v itemCfgId:%v", isWarehouse, v.ItemId, v.Class, v.Quality, v.Star, v.ItemValue, v.ItemQuality, v.ItemCfgId)
	}
	//-------- 日志打印
	this.ResetBag(user, isWarehouse, allBagAndWearEquip, allNotEquipAndBadEquipItems, allLowEquip, nil)

	return nil
}

func (this *BagManager) ResetBag(user *objs.User, isWarehouse bool, allBagAndWearEquip model.ItemMarkSlice3, allNotEquipAndBadEquipItems model.ItemMarkSlice2, allLowEquip model.ItemMarkSlice2, allBagSuperposition map[int]*model.ItemMark) model.Bag {

	if isWarehouse {
		// 更换位置
		pos := 0
		user.WarehouseBag = make(model.Bag, 0)
		if allBagSuperposition != nil {
			for _, v := range allBagSuperposition {
				if v.ItemType != pb.ITEMTYPE_EQUIP {

					posCount := v.Count / v.CountLimit
					posRemCount := v.Count % v.CountLimit
					if posRemCount > 0 {
						posCount += 1
					}
					if posCount == 0 {
						posCount += 1
					}
					logger.Debug("ResetBag  v.ItemId:%v  v.Count:%v v.CountLimit:%v   posCount:%v  posRemCount:%v", v.ItemId, v.Count, v.CountLimit, posCount, posRemCount)
					for i := 1; i <= posCount; i++ {
						if posCount == 1 {
							user.WarehouseBag = append(user.WarehouseBag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
							pos++
							break
						} else {
							if i == posCount {
								if posRemCount > 0 {
									user.WarehouseBag = append(user.WarehouseBag, &model.Item{ItemId: v.ItemId, Count: posRemCount, Position: pos, EquipIndex: v.EquipIndex})
									pos++
									break
								}
							}
							user.WarehouseBag = append(user.WarehouseBag, &model.Item{ItemId: v.ItemId, Count: v.CountLimit, Position: pos, EquipIndex: v.EquipIndex})
							pos++
						}
					}
				} else {
					user.WarehouseBag = append(user.WarehouseBag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
					pos++
				}
			}
			return user.WarehouseBag
		}

		if allBagAndWearEquip != nil {
			for _, v := range allBagAndWearEquip {
				if v.Index == -1 {
					continue
				}
				logger.Debug("背包整理完后 重新往背包里增加数据 isWarehouse:%v pos:%v  itemId:%v  count:%v  EquipIndex:%v", isWarehouse, pos, v.ItemId, v.Count, v.EquipIndex)
				user.WarehouseBag = append(user.WarehouseBag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
				pos++
			}
		}

		if allNotEquipAndBadEquipItems != nil {
			for _, v := range allNotEquipAndBadEquipItems {
				if v.Index == -1 {
					continue
				}
				logger.Debug("背包整理完后 重新往背包里增加数据 isWarehouse:%v pos:%v itemId:%v  count:%v  EquipIndex:%v", isWarehouse, pos, v.ItemId, v.Count, v.EquipIndex)
				user.WarehouseBag = append(user.WarehouseBag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
				pos++
			}
		}
		if allLowEquip != nil {
			for _, v := range allLowEquip {
				if v.Index == -1 {
					continue
				}
				logger.Debug("背包整理完后 重新往背包里增加数据 isWarehouse:%v pos:%v itemId:%v  count:%v  EquipIndex:%v", isWarehouse, pos, v.ItemId, v.Count, v.EquipIndex)
				user.WarehouseBag = append(user.WarehouseBag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
				pos++
			}
		}

		logger.Debug("整理完背包后初始化空余格子 isWarehouse:%v userId:%v  maxNum:%v", isWarehouse, user.Id, user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum)
		for i := pos; i < user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum; i++ {
			user.WarehouseBag = append(user.WarehouseBag, &model.Item{Position: i, ItemId: 0, Count: 0})
		}
		return user.WarehouseBag
	} else {
		// 更换位置
		pos := 0
		user.Bag = make(model.Bag, 0)
		if allBagSuperposition != nil {
			for _, v := range allBagSuperposition {
				if v.ItemType != pb.ITEMTYPE_EQUIP {
					posCount := v.Count / v.CountLimit
					posRemCount := v.Count % v.CountLimit
					if posRemCount > 0 {
						posCount += 1
					}
					if posCount == 0 {
						posCount += 1
					}
					for i := 1; i <= posCount; i++ {
						if posCount == 1 {
							user.Bag = append(user.Bag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
							pos++
							break
						} else {
							if i == posCount {
								if posRemCount > 0 {
									user.Bag = append(user.Bag, &model.Item{ItemId: v.ItemId, Count: posRemCount, Position: pos, EquipIndex: v.EquipIndex})
									pos++
									break
								}
							}
							user.Bag = append(user.Bag, &model.Item{ItemId: v.ItemId, Count: v.CountLimit, Position: pos, EquipIndex: v.EquipIndex})
							pos++
						}
					}
				} else {
					user.Bag = append(user.Bag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
					pos++
				}
			}
			return user.Bag
		}

		if allBagAndWearEquip != nil {
			for _, v := range allBagAndWearEquip {
				if v.Index == -1 {
					logger.Debug("背包整理 v.Index == -1  isWarehouse:%v  itemId:%v", isWarehouse, v.ItemId)
					continue
				}
				logger.Debug("背包整理完后 重新往背包里增加数据 isWarehouse:%v pos:%v  itemId:%v  count:%v  EquipIndex:%v", isWarehouse, pos, v.ItemId, v.Count, v.EquipIndex)
				user.Bag = append(user.Bag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
				pos++
			}
		}

		if allNotEquipAndBadEquipItems != nil {
			for _, v := range allNotEquipAndBadEquipItems {
				if v.Index == -1 {
					continue
				}
				logger.Debug("背包整理完后 重新往背包里增加数据 isWarehouse:%v pos:%v itemId:%v  count:%v  EquipIndex:%v", isWarehouse, pos, v.ItemId, v.Count, v.EquipIndex)
				user.Bag = append(user.Bag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
				pos++
			}
		}
		if allLowEquip != nil {
			for _, v := range allLowEquip {
				if v.Index == -1 {
					continue
				}
				logger.Debug("背包整理完后 重新往背包里增加数据 isWarehouse:%v pos:%v itemId:%v  count:%v  EquipIndex:%v", isWarehouse, pos, v.ItemId, v.Count, v.EquipIndex)
				user.Bag = append(user.Bag, &model.Item{ItemId: v.ItemId, Count: v.Count, Position: pos, EquipIndex: v.EquipIndex})
				pos++
			}
		}
		logger.Debug("整理完背包后初始化空余格子 userId:%v  maxNum:%v", user.Id, user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum)
		for i := pos; i < user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum; i++ {
			user.Bag = append(user.Bag, &model.Item{Position: i, ItemId: 0, Count: 0})
		}
		return user.Bag
	}
}
