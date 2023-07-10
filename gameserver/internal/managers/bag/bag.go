package bag

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewBagManager(module managersI.IModule) *BagManager {
	return &BagManager{IModule: module}
}

type BagManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *BagManager) Online(user *objs.User) {

	if user.BagInfo == nil {
		user.BagInfo = make(map[int]*model.BagInfoUnit)
	}

	//初始化玩家背包格数
	if user.BagInfo[constBag.BAG_TYPE_COMMON] == nil {
		user.BagInfo[constBag.BAG_TYPE_COMMON] = &model.BagInfoUnit{
			MaxNum:   gamedb.GetBagInitNum(),
			SpaceAdd: make(model.IntKv),
			BuyNum:   0,
		}
	}

	emptyPos := make(map[int]bool)
	for i := 0; i < user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum; i++ {
		emptyPos[i] = true
	}
	//记录玩家已背包格子数据
	bagData := make(model.Bag, user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum)
	for _, v := range user.Bag {
		if v.Position+1 > user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum {
			continue
		}
		bagData[v.Position] = v
		delete(emptyPos, v.Position)
	}
	user.Bag = bagData
	//初始化玩家背包格子数据
	for k, _ := range emptyPos {
		user.Bag[k] = &model.Item{Position: k}
		bagItem(user.Bag[k], 0, 0)
	}
	day := common.GetResetTime(time.Now())
	if user.RedPacketItem.Day != day {
		user.RedPacketItem = &model.RedPacketItem{
			Day:      day,
			PickNum:  0,
			PickInfo: make(model.IntKv),
		}
	}
}

func (this *BagManager) GetItemByPosition(user *objs.User, position int) *model.Item {

	if position >= user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum {
		return nil
	}
	return user.Bag[position]
}

func (this *BagManager) GetItemSpaceNum(itemId, count int) int {
	itemConf := gamedb.GetItemBaseCfg(itemId)
	//计算需要位置数
	addNeedEmptyNum := 1
	if itemConf.CountLimit > 0 {
		r := count % itemConf.CountLimit
		addNeedEmptyNum = count / itemConf.CountLimit
		if r > 0 {
			addNeedEmptyNum += 1
		}
	}
	return addNeedEmptyNum
}

/**
 *  @Description: 检查是否有足够空间格子
 *  @param user
 *  @param items
 *  @return bool
 */
func (this *BagManager) CheckHasEnoughPos(user *objs.User, items []*gamedb.ItemInfo) bool {

	needPosition := 0
	duidieItem := make(map[int]int)
	for _, v := range items {

		itemConf := gamedb.GetItemBaseCfg(v.ItemId)
		if this.isInBag(v.ItemId) {
			continue
		}
		if itemConf.CountLimit == 1 {
			needPosition += v.Count
		} else {
			duidieItem[v.ItemId] += v.Count
		}
	}

	emptyPos := 0
	for _, v := range user.Bag {

		if v.ItemId == 0 {
			if needPosition > 0 {
				needPosition -= 1
			} else {
				emptyPos += 1
			}
		} else {
			if count, ok := duidieItem[v.ItemId]; ok {
				itemConf := gamedb.GetItemBaseCfg(v.ItemId)
				if itemConf.CountLimit == 0 {
					delete(duidieItem, v.ItemId)
				} else if itemConf.CountLimit > 1 {
					if v.Count+count <= itemConf.CountLimit {
						delete(duidieItem, v.ItemId)
					} else {
						duidieItem[v.ItemId] = itemConf.CountLimit - v.Count
					}
				}
			}
		}
		if needPosition == 0 && len(duidieItem) == 0 {
			return true
		}
	}

	if needPosition > 0 {
		return false
	}

	for k, v := range duidieItem {

		needPos := this.GetItemSpaceNum(k, v)
		emptyPos -= needPos
	}
	if emptyPos >= 0 {
		return true
	} else {
		return false
	}
}

//获取空位
func (this *BagManager) getEmptyPos(user *objs.User, needSpaceNum int) []int {

	empty := make([]int, 0)
	getNum := 0
	for _, v := range user.Bag {
		if v.ItemId == 0 {
			empty = append(empty, v.Position)
			getNum++
			if getNum >= needSpaceNum {
				break
			}
		}
	}
	return empty
}

func (this *BagManager) GetItemNum(user *objs.User, itemId int) (int, error) {

	itemT := gamedb.GetItemBaseCfg(itemId)
	if itemT == nil {
		logger.Error("道具配置不存在，itemID:%v", itemId)
		return -1, gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg(itemId)
	}
	if itemT.Type == pb.ITEMTYPE_TOP {
		return user.GetTopDataByItemId(itemId), nil
	}

	findItem := make(map[int]bool)
	findItem[itemId] = true
	sameGroupItemId := this.getSameGroupItemId(itemId)
	if sameGroupItemId > 0 {
		findItem[sameGroupItemId] = true
	}

	count := 0
	for _, v := range user.Bag {
		if findItem[v.ItemId] {
			count += v.Count
		}
	}
	return count, nil
}

/**
 *  @Description: 获取同组另一个itemId
 *  @param itemId
 *  @return int
 */
func (this *BagManager) getSameGroupItemId(itemId int) int {

	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf.BindGruonp > 0 {
		itemBindGroupConf := gamedb.GetBindGroupBindGroupCfg(itemConf.BindGruonp)
		for _, v := range itemBindGroupConf.Itemid {
			if v != itemId {
				return v
			}
		}
	}
	return -1
}

func (this *BagManager) HasEnough(user *objs.User, itemId int, count int) (flag bool, num int) {

	if itemId == pb.ITEMID_BACK_CITY || itemId == pb.ITEMID_RANDOM_STONE {
		return true, 1
	}

	hasNum, err := this.GetItemNum(user, itemId)
	if err != nil {
		return false, hasNum
	}
	return hasNum >= count, hasNum
}

func (this *BagManager) HasEnoughItems(user *objs.User, itemInfos gamedb.ItemInfos) (flag bool, hasNum map[int]int) {
	userBag := user.Bag
	findMap := make(map[int]int)
	flag = true
	hasNum = make(map[int]int)

	needMap := make(map[int]int)
	for _, info := range itemInfos {
		needMap[info.ItemId] += info.Count
	}
	for itemId := range needMap {
		switch itemId {
		case pb.ITEMID_BACK_CITY:
			fallthrough
		case pb.ITEMID_RANDOM_STONE:
			hasNum[itemId] = 1
		default:
			if itemT := gamedb.GetItemBaseCfg(itemId); itemT != nil {
				if itemT.Type == pb.ITEMTYPE_TOP {
					hasNum[itemId] = user.GetTopDataByItemId(itemId)
				} else {
					findMap[itemId] = 0
				}
			} else {
				hasNum[itemId] = 0
			}
		}
	}
	for _, item := range userBag {
		itemId := item.ItemId
		if _, ok := findMap[itemId]; ok {
			hasNum[itemId] += item.Count
		}
	}
	for itemId, count := range needMap {
		if count > hasNum[itemId] {
			flag = false
			break
		}
	}
	return
}

func (this *BagManager) checkBagPosition(user *objs.User, bagType int, position int) bool {

	return position < user.BagInfo[bagType].MaxNum
}

func bagItem(item *model.Item, itemId, count int) {
	item.ItemId = itemId
	item.Count = count
	item.EquipIndex = 0
	//itemConf := gamedb.GetItemBaseCfg(itemId)
	//if itemConf != nil && (itemConf.Type == pb.ITEMTYPE_EQUIP || itemConf.Type == pb.ITEMTYPE_ZODIAC || itemConf.Type == pb.ITEMTYPE_KINGARMS || itemConf.Type == pb.ITEMTYPE_DRAGONARMS) {
	//	if source != nil {
	//		item.GetSource = source
	//	} else {
	//		item.GetSource = &model.GetSource{
	//			SkillUser: gamedb.SYSTEMADD,
	//			SkillDate: common.GetFormatTime(),
	//		}
	//	}
	//}
}

/**
 *  @Description: 重置数据，不重置位置
 *  @param item
 */
func BagItemUnitReset(item *model.Item) {
	item.ItemId = 0
	item.Count = 0
	item.EquipIndex = 0
}

func (this *BagManager) GetEquipItemInfos(user *objs.User, itemId int) (int, int, int) {

	for _, data := range user.Bag {
		if data.ItemId == itemId {
			return data.Position, data.ItemId, data.EquipIndex
		}
	}
	return -1, -1, -1
}

/**
*  @Description: 检查道具是否进背包
*  @receiver this
*  @param itemId
*  @return bool
**/
func (this *BagManager) isInBag(itemId int) bool {
	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf.Type == pb.ITEMTYPE_TOP || itemConf.Type == pb.ITEMTYPE_CONTRIBUTION {
		return true
	}
	return false
}
