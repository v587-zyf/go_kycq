package compose

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 传世装备合成
 *  @param user
 *  @param subId	合成id
 *  @param op
 *  @return error
 */
func (this *ComposeManager) ComposeChuanShiEquip(user *objs.User, subId int, op *ophelper.OpBagHelperDefault) error {
	if subId < 1 {
		return gamedb.ERRPARAM
	}
	chuanShiSubCfg := gamedb.GetComposeChuanShiSubComposeChuanShiSubCfg(subId)
	if chuanShiSubCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, -1, chuanShiSubCfg.ComCondition); !check {
		return gamedb.ERRCONDITION
	}
	bagConsume, heroConsume := make(gamedb.ItemInfos, 0), make(map[int]int)
	userHeros := user.Heros
	upDataCombat := false
	for _, info := range chuanShiSubCfg.Consume1 {
		needItemId, needCount := info.ItemId, info.Count
		hasEnough, hasNum := this.GetBag().HasEnough(user, needItemId, needCount)
		removeNum := needCount
		if !hasEnough {
			removeNum = hasNum
			if chuanshiCfg := gamedb.GetChuanShiEquipChuanShiEquipCfg(needItemId); chuanshiCfg != nil {
				for heroIndex, hero := range userHeros {
					equipPos := chuanshiCfg.Type
					if hero.ChuanShi[equipPos] == needItemId {
						heroConsume[heroIndex] = equipPos
						upDataCombat = true
						hasNum++
						if hasNum == needCount {
							break
						}
					}
				}
			}
			if hasNum < needCount {
				return gamedb.ERRNOTENOUGHGOODS
			}
		}
		bagConsume = append(bagConsume, &gamedb.ItemInfo{ItemId: needItemId, Count: removeNum})
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, bagConsume); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	for heroIndex, pos := range heroConsume {
		userHeros[heroIndex].ChuanShi[pos] = 0
		this.GetUserManager().SendMessage(user, &pb.ChuanShiRemoveAck{HeroIndex: int32(heroIndex), EquipPos: int32(pos)}, true)
	}
	addItems := make(gamedb.ItemInfos, 0)
	for _, ints := range chuanShiSubCfg.Composeid {
		addItems = append(addItems, &gamedb.ItemInfo{ItemId: ints[0], Count: ints[1]})
	}
	this.GetBag().AddItems(user, addItems, op)

	if upDataCombat {
		this.GetUserManager().UpdateCombat(user, -1)
	}
	user.Dirty = true
	return nil
}
