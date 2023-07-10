package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IBagManager interface {
	Online(user *objs.User)
	//增加道具
	Add(user *objs.User, helper *ophelper.OpBagHelperDefault, itemId, count int) error
	AddItem(user *objs.User, helper *ophelper.OpBagHelperDefault, itemId, count int) error
	/**
	 *  @Description:
	 *  @param user
	 *  @param items
	 *  @param helper
	 *  @return bool 是否进入邮件
	 */
	AddItems(user *objs.User, items gamedb.ItemInfos, helper *ophelper.OpBagHelperDefault) bool
	//移除道具
	Remove(user *objs.User, opHelper *ophelper.OpBagHelperDefault, itemId, count int) error
	//移除指定所有道具
	RemoveAllByItemId(user *objs.User, ophelper *ophelper.OpBagHelperDefault, itemId int) int
	//移除指定所有道具根据位置
	RemoveByPosition(user *objs.User, opHelper *ophelper.OpBagHelperDefault, itemId, count, pos int) error
	//移除ItemInfos
	RemoveItemsInfos(user *objs.User, op *ophelper.OpBagHelperDefault, itemInfos gamedb.ItemInfos) error
	//ka扩充背包
	BagSpaceAdd(user *objs.User, op *ophelper.OpBagHelperDefault) error
	//根据条件增加背包格子
	BagSpaceAddByType(user *objs.User, addType int) error
	//整理
	BagSort(user *objs.User, isWareHouse bool) error
	//是否有足够道具
	HasEnough(user *objs.User, itemId int, count int) (flag bool, num int)
	HasEnoughItems(user *objs.User, itemInfos gamedb.ItemInfos) (flag bool, hasNum map[int]int)
	//获取道具数量
	GetItemNum(user *objs.User, itemId int) (int, error)
	/**
    *  @Description: 检查背包位置是否足够
    *  @param user
    *  @param items
    *  @return bool
    **/
	CheckHasEnoughPos(user *objs.User, items []*gamedb.ItemInfo) bool
	/**
	 *  @Description: 装备替换
	 *  @param user	玩家数据
	 *  @param op	道具来源
	 *  @param equipBagPos 替换装备在背包的位置
	 *  @param equip	被替换的装备
	 */
	EquipChange(user *objs.User, op *ophelper.OpBagHelperDefault, equipBagPos int, equipTyp int, equip *model.Equip) (*model.Equip, error)
	/**
	 *  @Description: 通过位置获取道具信息
	 *  @param user
	 *  @param position
	 *  @return *model.Item
	 */
	GetItemByPosition(user *objs.User, position int) *model.Item
	/**
	 *  @Description:	装备锁定
	 *  @param user
	 *  @param position
	 */
	EquipLock(user *objs.User, position int) error

	/**
	 *  @Description: 装备回收
	 *  @param user
	 *  @param positions
	 */
	EquipRecover(user *objs.User, op *ophelper.OpBagHelperDefault, positions []int32) error

	/**
	 *  @Description: 		道具使用
	 *  @param user
	 *  @param itemId
	 *  @param itemNum
	 *  @param helperDefault
	 *  @return error
	 */
	ItemUse(user *objs.User, heroIndex, itemId, itemNum int, helperDefault *ophelper.OpBagHelperDefault) error

	//
	//  WareHouseOnline
	//  @Description:  仓库背包初始化检查
	//  @param user
	//
	WareHouseOnline(user *objs.User)
	//仓库扩容
	WareHouseBagSpaceAdd(user *objs.User, op *ophelper.OpBagHelperDefault) error
	//背包东西移动到仓库
	WareHouseBagAdd(user *objs.User, op *ophelper.OpBagHelperDefault, positions []int32) error
	//仓库东西移动到背包
	WareHouseMoveToBag(user *objs.User, op *ophelper.OpBagHelperDefault, positions []int32) error
	//销毁
	EquipDestroy(user *objs.User, op *ophelper.OpBagHelperDefault, positions, count int) error
	//gm清理
	BagClear(user *objs.User)
	//Position, ItemId, EquipIndex
	GetEquipItemInfos(user *objs.User, itemId int) (int, int, int)
}
