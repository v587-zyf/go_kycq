package equip

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
	"math"
)

const (
	STRENGTHEN_BREAK_NO = iota
	STRENGTHEN_BREAK_YES
)

var STRENGTHEN_POS = []int{
	pb.EQUIPPOS_ONE,   //主武器
	pb.EQUIPPOS_FOUR,  //衣服
	pb.EQUIPPOS_EIGHT, //手镯1
	pb.EQUIPPOS_SEVEN, //戒指1
	pb.EQUIPPOS_FIVE,  //腰带
	pb.EQUIPPOS_THREE, //头盔
	pb.EQUIPPOS_TWO,   //项链
	pb.EQUIPPOS_TEN,   //手镯2
	pb.EQUIPPOS_NINE,  //戒指2
	pb.EQUIPPOS_SIX,   //鞋子
}

/**
 *  @Description: 装备强化
 *  @param user
 *  @param heroIndex
 *  @param pos	装备位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *EquipManager) EquipsStrengthen(user *objs.User, heroIndex int, pos int, op *ophelper.OpBagHelperDefault, ack *pb.EquipStrengthenAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if !pb.EQUIPPOS_MAP[pos] {
		return gamedb.ERRNOTEQUIP
	}
	if hero.Equips[pos].ItemId == 0 {
		return gamedb.ERRNOTWEAREQUIP
	}

	lv := hero.EquipsStrength[pos]
	if lv >= gamedb.GetMaxValById(pos, constMax.MAX_EQUIP_STRENGTHEN_LEVEL) {
		return gamedb.ERREQUIPENOUGH
	}

	strengthenCfg := gamedb.GetEquipStrengthConfByLvAndPos(pos, lv)
	if strengthenCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, heroIndex, strengthenCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, strengthenCfg.Consume); err != nil {
		return err
	}

	beforeLv := lv
	randRes := common.RandByTenShousand(strengthenCfg.Rate)
	if randRes {
		lv++
		hero.EquipsStrength[pos] = lv
		this.GetTask().AddTaskProcess(user, pb.CONDITION_ALL_HEROS_QIANG_HUA_LV, -1)
	}
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HEROS_QIANG_HUA_LV, []int{})
	this.GetCondition().RecordCondition(user, pb.CONDITION_ONE_EQUIP_INTENSIFY_MAX_LV, []int{lv})
	ack.HeroIndex = int32(heroIndex)
	ack.EquipGrids = &pb.EquipGrid{Pos: int32(pos), Strength: int32(lv)}
	ack.IsUp = randRes
	kyEvent.EquipIntensify(user, heroIndex, pos, hero.Equips[pos].ItemId, beforeLv, hero.EquipsStrength[pos])
	this.GetUserManager().UpdateCombat(user, heroIndex)

	this.GetTask().AddTaskProcess(user, pb.CONDITION_QIANGHUA_EQUIP, 1)
	return nil
}

/**
 *  @Description: 装备一键强化、突破
 *  @param user
 *  @param heroIndex
 *  @param isBreak	是否突破
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *EquipManager) EquipStrengthenOneKey(user *objs.User, heroIndex int, isBreak bool, op *ophelper.OpBagHelperDefault, ack *pb.EquipStrengthenAutoAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	heroEquips, heroStrengthen := hero.Equips, hero.EquipsStrength
	removeItemInfos, removeMap := make(gamedb.ItemInfos, 0), make(map[int]int)

	okNum, maxNum := 0, 0
	minPos, minLv := 0, math.MaxInt32
	for _, pos := range STRENGTHEN_POS {
		if heroEquips[pos].ItemId < 1 {
			continue
		}
		if heroStrengthen[pos] < minLv {
			minPos = pos
			minLv = heroStrengthen[pos]
		}
	}
	if minLv >= gamedb.GetMaxValById(minPos, constMax.MAX_EQUIP_STRENGTHEN_LEVEL) {
		return gamedb.ERRLVENOUGH
	}
	for _, pos := range STRENGTHEN_POS {
		if heroEquips[pos].ItemId < 1 {
			continue
		}
		if heroStrengthen[pos] == minLv {
			strengthenCfg := gamedb.GetEquipStrengthConfByLvAndPos(pos, heroStrengthen[pos])
			if strengthenCfg == nil {
				continue
			}
			if check := this.GetCondition().CheckMulti(user, heroIndex, strengthenCfg.Condition); !check {
				continue
			}
			if !isBreak && strengthenCfg.IsBreak == STRENGTHEN_BREAK_YES {
				continue
			}

			for i := 0; i < constConstant.COMPUTE_TEN_THOUSAND; i++ {
				removeItemInfos = append(removeItemInfos, strengthenCfg.Consume...)
				if enough, _ := this.GetBag().HasEnoughItems(user, removeItemInfos); !enough {
					break
				}
				for _, info := range strengthenCfg.Consume {
					removeMap[info.ItemId] += info.Count
				}
				randRes := common.RandByTenShousand(strengthenCfg.Rate)
				if randRes {
					okNum++
					heroStrengthen[pos]++
					if maxNum < heroStrengthen[minPos] {
						maxNum = heroStrengthen[minPos]
					}
					break
				}
			}
		}
	}
	//okNum, maxNum := 0, 0
	//for i := 0; i < 50000; i++ {
	//	minPos, minLv := 0, math.MaxInt32
	//	for _, pos := range STRENGTHEN_POS {
	//		if heroEquips[pos].ItemId < 1 {
	//			continue
	//		}
	//		if heroStrengthen[pos] < minLv {
	//			minPos = pos
	//			minLv = heroStrengthen[pos]
	//		}
	//	}
	//	if minPos == 0 || minLv >= gamedb.GetMaxValById(minPos, constMax.MAX_EQUIP_STRENGTHEN_LEVEL) {
	//		break
	//	}
	//	strengthenCfg := gamedb.GetEquipStrengthConfByLvAndPos(minPos, minLv)
	//	if strengthenCfg == nil {
	//		break
	//	}
	//	if check := this.GetCondition().CheckMulti(user, heroIndex, strengthenCfg.Condition); !check {
	//		break
	//	}
	//	if !isBreak && strengthenCfg.IsBreak == STRENGTHEN_BREAK_YES {
	//		break
	//	}
	//
	//	enough := true
	//	for i := 0; i < 100; i++ {
	//		removeItemInfos = append(removeItemInfos, strengthenCfg.Consume...)
	//		enough, _ = this.GetBag().HasEnoughItems(user, removeItemInfos)
	//		if !enough {
	//			break
	//		}
	//		for _, info := range strengthenCfg.Consume {
	//			removeMap[info.ItemId] += info.Count
	//		}
	//		randRes := common.RandByTenShousand(strengthenCfg.Rate)
	//		if randRes {
	//			okNum++
	//			heroStrengthen[minPos]++
	//			if maxNum < heroStrengthen[minPos] {
	//				maxNum = heroStrengthen[minPos]
	//			}
	//			break
	//		}
	//	}
	//	if !enough {
	//		break
	//	}
	//}
	//if len(removeMap) > 0 {
	//	removeItems := make(gamedb.ItemInfos, 0)
	//	for itemId, count := range removeMap {
	//		removeItems = append(removeItems, &gamedb.ItemInfo{ItemId: itemId, Count: count})
	//	}
	//	this.GetBag().RemoveItemsInfos(user, op, removeItems)
	//}
	if len(removeMap) > 0 {
		removeItems := make(gamedb.ItemInfos, 0)
		for itemId, count := range removeMap {
			removeItems = append(removeItems, &gamedb.ItemInfo{ItemId: itemId, Count: count})
		}
		if err := this.GetBag().RemoveItemsInfos(user, op, removeItems); err != nil {
			return err
		}
	}

	this.GetCondition().RecordCondition(user, pb.CONDITION_ONE_EQUIP_INTENSIFY_MAX_LV, []int{maxNum})
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_ALL_HEROS_QIANG_HUA_LV, okNum)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_QIANGHUA_EQUIP, okNum)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HEROS_QIANG_HUA_LV, []int{})

	ack.HeroIndex = int32(heroIndex)
	ack.EquipGrids = builder.BuilderEquipStrength(hero)
	return nil
}
