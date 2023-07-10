package compose

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strings"
)

func NewComposeManager(module managersI.IModule) *ComposeManager {
	return &ComposeManager{IModule: module}
}

type ComposeManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 合成道具
 *  @param user
 *  @param heroIndex
 *  @param subId		合成道具表id
 *  @param composeNum	合成数量
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *ComposeManager) Compose(user *objs.User, heroIndex, subId, composeNum int, op *ophelper.OpBagHelperDefault, ack *pb.ComposeAck) error {
	if heroIndex <= 0 || subId <= 0 || composeNum <= 0 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	subConf := gamedb.GetComposeSubComposeSubCfg(subId)
	if subConf == nil {
		return gamedb.ERRPARAM
	}

	composeId, composeCount := subConf.Composeid.ItemId, subConf.Composeid.Count*composeNum

	consumes := make(gamedb.ItemInfos, 0)
	var checkItemFunc = func(consume gamedb.ItemInfos, composeNums, coinConsumeNum int) (err error, needItemId, needConsumeCount, needItemNum int) {
		for _, itemInfo := range consume {
			consumeItemId, consumeCount := itemInfo.ItemId, itemInfo.Count*composeNums
			needItemCfg := gamedb.GetItemBaseCfg(consumeItemId)
			if needItemCfg.Type != pb.ITEMTYPE_TOP {
				needItemId = consumeItemId
				needConsumeCount = consumeCount
			}
			if enough, hasNum := this.GetBag().HasEnough(user, consumeItemId, consumeCount); !enough {
				if needItemCfg.Type == pb.ITEMTYPE_TOP {
					err = gamedb.ERRNOTENOUGHGOODS
					return
				}
				consumes = append(consumes, &gamedb.ItemInfo{
					ItemId: consumeItemId,
					Count:  hasNum,
				})
				needItemNum = consumeCount - hasNum
			} else {
				removeCount := consumeCount
				//if needItemCfg.Type == pb.ITEMTYPE_TOP {
				//	removeCount = itemInfo.Count * coinConsumeNum
				//}
				consumes = append(consumes, &gamedb.ItemInfo{
					ItemId: consumeItemId,
					Count:  removeCount,
				})
			}
		}
		return
	}

	var err error
	var enough bool
	var needItemId, needConsumeCount, needItemNum, coinConsumeNum int
	for i := 0; i < 10000; i++ {
		if needItemId == 0 {
			coinConsumeNum = composeNum
			err, needItemId, needConsumeCount, needItemNum = checkItemFunc(subConf.Consume1, composeNum, composeNum)
		} else {
			composeItemCfg := gamedb.GetComposeItemCfg(needItemId)
			if composeItemCfg == nil {
				break
			}
			coinConsumeNum *= needConsumeCount
			err, needItemId, needConsumeCount, needItemNum = checkItemFunc(composeItemCfg.Consume1, needItemNum, coinConsumeNum)
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
	if err = this.GetBag().RemoveItemsInfos(user, op, consumes); err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}

	this.GetBag().Add(user, op, composeId, composeCount)
	itemInfo := gamedb.GetItemBaseCfg(composeId)
	if strings.Contains(itemInfo.Name, "灵丹") {
		this.GetCondition().RecordCondition(user, pb.CONDITION_HE_CHENG_LIN_DAN, []int{composeCount})
		this.GetTask().AddTaskProcess(user, pb.CONDITION_HE_CHENG_LIN_DAN, -1)
	}
	user.Dirty = true
	kyEvent.Compose(user, composeId, composeCount, consumes)
	this.GetUserManager().UpdateCombat(user, -1)
	ack.HeroIndex = int32(heroIndex)
	ack.Goods = op.ToChangeItems()
	return nil
}
