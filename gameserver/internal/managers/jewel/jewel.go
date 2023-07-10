package jewel

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strconv"
)

var jewelMakePosSlice = []int{1, 3, 4, 2, 8, 10, 7, 9, 5, 6}

func NewJewelManager(m managersI.IModule) *JewelManager {
	return &JewelManager{
		IModule: m,
	}
}

type JewelManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 宝石一件穿戴
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *JewelManager) JewelMakeAll(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.JewelMakeAllAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}

	jewelItemIdMap := gamedb.GetJewelItemIdMap()
	bagHasJewels := make(map[int]int)
	bag := user.Bag
	for _, item := range bag {
		if _, ok := jewelItemIdMap[item.ItemId]; !ok {
			continue
		}
		bagHasJewels[item.ItemId] += item.Count
	}
	removeMap, addMap := make(map[int]int), make(map[int]int)
	for _, pos := range jewelMakePosSlice {
		if hero.Jewel[pos] == nil {
			hero.Jewel[pos] = &model.Jewel{}
		}
		heroJewel := hero.Jewel[pos]
		jewelBodyCfg := gamedb.GetJewelBodyJewelBodyCfg(pos)
		if jewelBodyCfg == nil {
			return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("jewelBody body " + strconv.Itoa(pos))
		}
		jewelSlice := gamedb.GetJewelTypeLvCfg(gamedb.GetJewelBodyJewelBodyCfg(pos).Type)
		if check := this.GetCondition().CheckMulti(user, heroIndex, jewelBodyCfg.Condition1); check {
			this.SetJewel(jewelSlice, bagHasJewels, removeMap, addMap, heroJewel, pb.JEWELPOS_ONE)
		}
		if check := this.GetCondition().CheckMulti(user, heroIndex, jewelBodyCfg.Condition2); check {
			this.SetJewel(jewelSlice, bagHasJewels, removeMap, addMap, heroJewel, pb.JEWELPOS_TWO)
		}
		if check := this.GetCondition().CheckMulti(user, heroIndex, jewelBodyCfg.Condition3); check {
			this.SetJewel(jewelSlice, bagHasJewels, removeMap, addMap, heroJewel, pb.JEWELPOS_THREE)
		}
	}
	if len(addMap) > 0 {
		addItems := make(gamedb.ItemInfos, 0)
		for itemId, count := range addMap {
			if count <= 0 {
				continue
			}
			addItems = append(addItems, &gamedb.ItemInfo{
				ItemId: itemId,
				Count:  count,
			})
		}
		this.GetBag().AddItems(user, addItems, op)
	}
	if len(removeMap) > 0 {
		removeItems := make(gamedb.ItemInfos, 0)
		for itemId, count := range removeMap {
			if count <= 0 {
				continue
			}
			removeItems = append(removeItems, &gamedb.ItemInfo{
				ItemId: itemId,
				Count:  count,
			})
		}
		if err := this.GetBag().RemoveItemsInfos(user, op, removeItems); err != nil {
			return gamedb.ERRNOTENOUGHGOODS
		}
	}

	ack.HeroIndex = int32(heroIndex)
	ack.Jewels = builder.BuilderJewels(hero)

	this.ChangeOperation(user, heroIndex)
	if len(removeMap) > 0 {
		this.GetCondition().RecordCondition(user, pb.CONDITION_XIANG_QIAN_BAO_SHI, []int{1})
		this.GetTask().AddTaskProcess(user, pb.CONDITION_XIANG_QIAN_BAO_SHI, -1)
		this.GetCondition().RecordCondition(user, pb.CONDITION_XIANG_QIAN_BAO_SHI_1, []int{})
	}
	return nil
}

func (this *JewelManager) SetJewel(jewelSlice []int, bagHasJewels, removeMap, addMap map[int]int, heroJewel *model.Jewel, pos int) {
	for _, itemId := range jewelSlice {
		if num, ok := bagHasJewels[itemId]; ok && num > 0 {
			jewelId, isRemove := 0, false
			switch pos {
			case pb.JEWELPOS_ONE:
				if heroJewel.One == 0 {
					isRemove = true
					heroJewel.One = itemId
				} else {
					jewelId = heroJewel.One
				}
			case pb.JEWELPOS_TWO:
				if heroJewel.Two == 0 {
					isRemove = true
					heroJewel.Two = itemId
				} else {
					jewelId = heroJewel.Two
				}
			case pb.JEWELPOS_THREE:
				if heroJewel.Three == 0 {
					isRemove = true
					heroJewel.Three = itemId
				} else {
					jewelId = heroJewel.Three
				}
			}
			if jewelId != 0 {
				jewelCfg := gamedb.GetJewelJewelCfg(jewelId)
				itemJewelCfg := gamedb.GetJewelJewelCfg(itemId)
				if jewelCfg.Level < itemJewelCfg.Level {
					switch pos {
					case pb.JEWELPOS_ONE:
						heroJewel.One = itemId
					case pb.JEWELPOS_TWO:
						heroJewel.Two = itemId
					case pb.JEWELPOS_THREE:
						heroJewel.Three = itemId
					}
					isRemove = true
					addMap[jewelId]++
					bagHasJewels[jewelId]++
				}
			}
			if isRemove {
				if hasNum, ok := addMap[itemId]; ok && hasNum > 0 {
					addMap[itemId]--
				} else {
					removeMap[itemId]++
				}
				bagHasJewels[itemId]--
				break
			}
		}
	}
}

/**
 *  @Description: 宝石穿戴一个部位
 *  @param user
 *  @param heroIndex
 *  @param equipPos	装备部位
 *  @param jewelPos	宝石部位
 *  @param itemId	宝石id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *JewelManager) JewelMake(user *objs.User, heroIndex, equipPos, jewelPos, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.JewelMakeAck) error {
	if itemId < 1 {
		return gamedb.ERRPARAM
	}
	hero, err := this.checkArgs(user, heroIndex, equipPos, jewelPos)
	if err != nil {
		return err
	}

	heroJewel := hero.Jewel[equipPos]
	if heroJewel == nil {
		hero.Jewel[equipPos] = &model.Jewel{}
		heroJewel = hero.Jewel[equipPos]
	}
	jewelBodyCfg := gamedb.GetJewelBodyJewelBodyCfg(equipPos)
	if jewelBodyCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("jewelBody body " + strconv.Itoa(equipPos))
	}
	condition := make(map[int]int)
	var jewelId int
	switch jewelPos {
	case pb.JEWELPOS_ONE:
		jewelId = heroJewel.One
		condition = jewelBodyCfg.Condition1
	case pb.JEWELPOS_TWO:
		jewelId = heroJewel.Two
		condition = jewelBodyCfg.Condition2
	case pb.JEWELPOS_THREE:
		jewelId = heroJewel.Three
		condition = jewelBodyCfg.Condition3
	}
	if jewelId > 0 {
		return gamedb.ERRREPEATACTIVE
	}
	if check := this.GetCondition().CheckMulti(user, heroIndex, condition); !check {
		return gamedb.ERRCONDITION
	}

	if err := this.GetBag().Remove(user, op, itemId, 1); err != nil {
		return err
	}

	this.UpdateJewelData(heroJewel, jewelPos, itemId)

	ack.HeroIndex = int32(heroIndex)
	ack.EquipPos = int32(equipPos)
	ack.JewelPos = int32(jewelPos)
	ack.Jewel = builder.BuildJewel(heroJewel)
	kyEvent.JewelMake(user, heroIndex, equipPos, jewelPos, itemId)

	this.ChangeOperation(user, heroIndex)
	this.GetCondition().RecordCondition(user, pb.CONDITION_XIANG_QIAN_BAO_SHI, []int{1})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_XIANG_QIAN_BAO_SHI, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_XIANG_QIAN_BAO_SHI_1, []int{})
	return nil
}

/**
 *  @Description: 宝石升级
 *  @param user
 *  @param heroIndex
 *  @param equipPos	装备部位
 *  @param jewelPos	宝石部位
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *JewelManager) JewelUpLv(user *objs.User, heroIndex, equipPos, jewelPos int, op *ophelper.OpBagHelperDefault, ack *pb.JewelUpLvAck) error {
	hero, err := this.checkArgs(user, heroIndex, equipPos, jewelPos)
	if err != nil {
		return err
	}

	heroJewel := hero.Jewel[equipPos]
	if heroJewel == nil {
		return gamedb.ERRNOTACTIVE
	}

	var jewelId int
	switch jewelPos {
	case pb.JEWELPOS_ONE:
		jewelId = heroJewel.One
	case pb.JEWELPOS_TWO:
		jewelId = heroJewel.Two
	case pb.JEWELPOS_THREE:
		jewelId = heroJewel.Three
	}
	if jewelId == 0 {
		return gamedb.ERRNOTACTIVE
	}
	jewelCfg := gamedb.GetJewelJewelCfg(jewelId)
	jewelMaxLv := gamedb.GetMaxValById(jewelCfg.Type, constMax.MAX_JEWEL_LEVEL)
	if jewelCfg.Level >= jewelMaxLv {
		return gamedb.ERRLVENOUGH
	}

	removeMap := make(map[int]int)
	var checkFunc = func(needItem map[int]int) (err error, needItemId, needItemNum int) {
		for itemId, itemCount := range needItem {
			itemCfg := gamedb.GetItemBaseCfg(itemId)
			if itemCfg == nil {
				logger.Warn("JewelUpLv item not found itemId:%v", itemId)
				continue
			}
			enough, hasNum := this.GetBag().HasEnough(user, itemId, itemCount)
			if !enough {
				if itemCfg.Type == pb.ITEMTYPE_TOP {
					err = gamedb.ERRNOTENOUGHGOODS
					return
				}
				removeMap[itemId] += hasNum
				needItemId = itemId
				needItemNum = itemCount - hasNum
			} else {
				removeMap[itemId] += itemCount
			}
		}
		return
	}
	var enough bool
	var needItemId, needItemNum int
	for i := 0; i < 10000; i++ {
		if needItemId == 0 {
			needMap := make(map[int]int)
			for _, itemInfo := range jewelCfg.Consume {
				needMap[itemInfo.ItemId] += itemInfo.Count
			}
			err, needItemId, needItemNum = checkFunc(needMap)
		} else {
			jewelCfg := gamedb.GetJewelJewelCfg(needItemId)
			composeItemCfg := gamedb.GetJewelByKindAndLv(jewelCfg.Type, jewelCfg.Level-1)
			if composeItemCfg == nil {
				break
			}
			needMap := make(map[int]int)
			for _, itemInfo := range composeItemCfg.Consume {
				needMap[itemInfo.ItemId] += itemInfo.Count * needItemNum
			}
			err, needItemId, needItemNum = checkFunc(needMap)
		}
		if err != nil {
			return err
		}
		if needItemNum == 0 {
			enough = true
			break
		}
	}
	if !enough {
		return gamedb.ERRNOTENOUGHGOODS
	}
	if len(removeMap) > 0 {
		removeInfos := make(gamedb.ItemInfos, 0)
		for itemId, count := range removeMap {
			removeInfos = append(removeInfos, &gamedb.ItemInfo{ItemId: itemId, Count: count})
		}
		if err := this.GetBag().RemoveItemsInfos(user, op, removeInfos); err != nil {
			return err
		}
	}

	jewelNextCfg := gamedb.GetJewelByKindAndLv(jewelCfg.Type, jewelCfg.Level+1)
	this.UpdateJewelData(heroJewel, jewelPos, jewelNextCfg.Id)

	ack.HeroIndex = int32(heroIndex)
	ack.EquipPos = int32(equipPos)
	ack.JewelPos = int32(jewelPos)
	ack.Jewel = builder.BuildJewel(heroJewel)

	this.ChangeOperation(user, heroIndex)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_XIANG_QIAN_BAO_SHI_1, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_XIANG_QIAN_BAO_SHI_1, []int{})
	return nil
}

/**
 *  @Description: 宝石安装
 *  @param user
 *  @param heroIndex
 *  @param equipPos 装备部为
 *  @param jewelPos	宝石部位
 *  @param itemId	宝石id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *JewelManager) JewelChange(user *objs.User, heroIndex, equipPos, jewelPos, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.JewelChangeAck) error {
	if itemId < 1 {
		return gamedb.ERRPARAM
	}
	hero, err := this.checkArgs(user, heroIndex, equipPos, jewelPos)
	if err != nil {
		return err
	}

	heroJewel := hero.Jewel[equipPos]
	if heroJewel == nil {
		return gamedb.ERRNOTACTIVE
	}
	jewelCfg := gamedb.GetJewelJewelCfg(itemId)
	if jewelCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("item" + strconv.Itoa(itemId))
	}

	if err := this.GetBag().Remove(user, op, itemId, 1); err != nil {
		return err
	}
	var jewelId int
	switch jewelPos {
	case pb.JEWELPOS_ONE:
		jewelId = heroJewel.One
	case pb.JEWELPOS_TWO:
		jewelId = heroJewel.Two
	case pb.JEWELPOS_THREE:
		jewelId = heroJewel.Three
	}
	if jewelId != 0 {
		if err := this.GetBag().Add(user, op, jewelId, 1); err != nil {
			return err
		}
	}
	this.UpdateJewelData(heroJewel, jewelPos, itemId)

	ack.HeroIndex = int32(heroIndex)
	ack.EquipPos = int32(equipPos)
	ack.JewelPos = int32(jewelPos)
	ack.Jewel = builder.BuildJewel(heroJewel)
	this.ChangeOperation(user, heroIndex)
	return nil
}

/**
 *  @Description: 宝石移除
 *  @param user
 *  @param heroIndex
 *  @param equipPos	装备部为
 *  @param jewelPos	宝石部位
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *JewelManager) JewelRemove(user *objs.User, heroIndex, equipPos, jewelPos int, op *ophelper.OpBagHelperDefault, ack *pb.JewelRemoveAck) error {
	hero, err := this.checkArgs(user, heroIndex, equipPos, jewelPos)
	if err != nil {
		return err
	}

	heroJewel := hero.Jewel[equipPos]
	if heroJewel == nil {
		return gamedb.ERRNOTACTIVE
	}

	var jewelId int
	switch jewelPos {
	case pb.JEWELPOS_ONE:
		jewelId = heroJewel.One
	case pb.JEWELPOS_TWO:
		jewelId = heroJewel.Two
	case pb.JEWELPOS_THREE:
		jewelId = heroJewel.Three
	}
	if jewelId == 0 {
		return gamedb.ERRPARAM
	}
	if err := this.GetBag().Add(user, op, jewelId, 1); err != nil {
		return err
	}
	this.UpdateJewelData(heroJewel, jewelPos, 0)

	ack.HeroIndex = int32(heroIndex)
	ack.EquipPos = int32(equipPos)
	ack.JewelPos = int32(jewelPos)
	ack.Jewel = builder.BuildJewel(heroJewel)
	this.ChangeOperation(user, heroIndex)
	return nil
}

func (this *JewelManager) UpdateJewelData(heroJewel *model.Jewel, jewelPos, val int) {
	switch jewelPos {
	case pb.JEWELPOS_ONE:
		heroJewel.One = val
	case pb.JEWELPOS_TWO:
		heroJewel.Two = val
	case pb.JEWELPOS_THREE:
		heroJewel.Three = val
	}
}

func (this *JewelManager) ChangeOperation(user *objs.User, heroIndex int) {
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.setMaxBaoSiLv(user)
}

func (this *JewelManager) checkArgs(user *objs.User, heroIndex int, equipPos int, jewelPos int) (*objs.Hero, error) {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return nil, gamedb.ERRHERONOTFOUND
	}
	if _, ok := pb.EQUIPPOS_MAP[equipPos]; !ok {
		return nil, gamedb.ERRPARAM
	}
	if _, ok := pb.JEWELPOS_MAP[jewelPos]; !ok {
		return nil, gamedb.ERRPARAM
	}
	return hero, nil
}

func (this *JewelManager) setMaxBaoSiLv(user *objs.User) {
	maxNum, _ := this.GetCondition().Check(user, -1, pb.CONDITION_XIANG_QIAN_BAO_SHI_1, 1)
	if maxNum > user.ModuleUpMax.BaoSiLv {
		user.ModuleUpMax.BaoSiLv = maxNum
	}
	user.Dirty = true
	return
}
