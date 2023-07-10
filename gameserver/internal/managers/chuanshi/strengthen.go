package chuanshi

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
)

var useStonePos = map[int]int{
	pb.CHUANSHIPOS_WU_QI: 0,
	pb.CHUANSHIPOS_YI_FU: 0,
}

/**
 *  @Description: 传世装备强化
 *  @param user
 *  @param heroIndex
 *  @param pos
 *  @param stone
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *ChuanShiManager) Strengthen(user *objs.User, heroIndex, pos, stone int, op *ophelper.OpBagHelperDefault, ack *pb.ChuanshiStrengthenAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	//if _, ok := useStonePos[pos]; stone > 0 && !ok {
	//	return gamedb.ERRPARAM
	//}
	equips := hero.ChuanShi
	if id, ok := equips[pos]; !ok || id < 1 {
		return gamedb.ERRNOTWEAREQUIP
	}
	strengthen := hero.ChuanshiStrengthen
	if gamedb.GetChuanShiStrengthenByPosAndLv(pos, strengthen[pos]+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	strengthenCfg := gamedb.GetChuanShiStrengthenByPosAndLv(pos, strengthen[pos])
	if len(strengthenCfg.Condition) > 0 {
		for cId, cVal := range strengthenCfg.Condition {
			if _, check := this.GetCondition().CheckBySlice(user, heroIndex, []int{cId, cVal}); !check {
				return gamedb.ERRCONDITION
			}
		}
	}
	rate := strengthenCfg.Rate
	removeMap := make(map[int]int)
	switch stone {
	case pb.CHUANSHISTRENGTHENSTONE_LUCK:
		itemId := 0
		if len(strengthenCfg.LuckyStone) < 1 {
			return gamedb.ERRITEMCANNOTUSE
		}
		for _, info := range strengthenCfg.LuckyStone {
			removeMap[info.ItemId] += info.Count
			itemId = info.ItemId
		}
		rate += gamedb.GetItemBaseCfg(itemId).EffectVal
	case pb.CHUANSHISTRENGTHENSTONE_PROTECT:
		if len(strengthenCfg.Relegation) < 1 {
			return gamedb.ERRITEMCANNOTUSE
		}
		for _, info := range strengthenCfg.Relegation {
			removeMap[info.ItemId] += info.Count
		}
	case pb.CHUANSHISTRENGTHENSTONE_RISE:
		if len(strengthenCfg.Definitely) < 1 {
			return gamedb.ERRITEMCANNOTUSE
		}
		for _, info := range strengthenCfg.Definitely {
			removeMap[info.ItemId] += info.Count
		}
		rate += constConstant.COMPUTE_TEN_THOUSAND
	}
	for _, info := range strengthenCfg.Consume {
		removeMap[info.ItemId] += info.Count
	}
	if len(removeMap) > 0 {
		removeItems := make(gamedb.ItemInfos, 0)
		for itemId, count := range removeMap {
			removeItems = append(removeItems, &gamedb.ItemInfo{
				ItemId: itemId,
				Count:  count,
			})
		}
		if err := this.GetBag().RemoveItemsInfos(user, op, removeItems); err != nil {
			return err
		}
	}

	isUp := common.RandByTenShousand(rate)
	if isUp {
		strengthen[pos]++
	} else {
		if stone != pb.CHUANSHISTRENGTHENSTONE_PROTECT {
			if strengthen[pos]-1 >= 0 && common.RandByTenShousand(strengthenCfg.DemoteRate) {
				strengthen[pos]--
			}
		}
	}
	this.GetUserManager().UpdateCombat(user, heroIndex)
	user.Dirty = true

	ack.HeroIndex = int32(heroIndex)
	ack.EquipPos = int32(pos)
	ack.Lv = int32(strengthen[pos])
	ack.IsUp = isUp
	this.GetTask().UpdateTaskProcess(user, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_CHUAN_SI_QIANG_HUA_ALL_LV, []int{})
	return nil
}
