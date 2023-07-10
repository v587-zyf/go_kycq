package bag

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

func (this *BagManager) BagSpaceAdd(user *objs.User, op *ophelper.OpBagHelperDefault) error {

	buyNum := user.BagInfo[constBag.BAG_TYPE_COMMON].BuyNum
	if buyNum >= gamedb.GetBagAddCfgLen(constBag.BAG_ADD_TYPE_ITEM) {
		return gamedb.ERRCANNOTBUYSPACE
	}
	costItemId, costCount, addNum := gamedb.GetBagSpaceAddCost(buyNum)
	hasEnouth, _ := this.HasEnough(user, costItemId, costCount)
	if !hasEnouth {
		return gamedb.ERRNOTENOUGHGOODS
	}
	err := this.Remove(user, op, costItemId, costCount)
	if err != nil {
		return err
	}
	this.bagSpaceAdd(user, addNum)
	return nil
}

func (this *BagManager) bagSpaceAdd(user *objs.User, addNum int) {
	nowNum := user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum
	maxNum := nowNum + addNum
	user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum = maxNum
	for ; nowNum < maxNum; nowNum++ {
		user.Bag = append(user.Bag, &model.Item{Position: nowNum})
	}
	user.BagInfo[constBag.BAG_TYPE_COMMON].BuyNum += 1
}

func (this *BagManager) BagSpaceAddByType(user *objs.User, addType int) error {

	getValue := 0
	nowValue := 0
	if addType == constBag.BAG_ADD_TYPE_VIP {
		nowValue = user.VipLevel
		getValue = user.BagInfo[constBag.BAG_TYPE_COMMON].SpaceAdd[constBag.BAG_ADD_TYPE_VIP]

	} else if addType == constBag.BAG_ADD_TYPE_ONLINE {
		//TODO 在线时长增加背包格子
		return nil
	} else {
		return gamedb.ERRPARAM
	}
	//背包根据vip增加格子
	newGetValue, addNum := gamedb.GetBagSapceAddByType(addType, getValue, nowValue)
	if addNum == 0 {
		return nil
	}
	this.bagSpaceAdd(user, addNum)
	user.BagInfo[constBag.BAG_TYPE_COMMON].SpaceAdd[addType] = newGetValue
	logger.Debug("根据条件增加背包格子，玩家：%v-%v,类型：%v,已获取：%v，当前值：%v,增加：%v,新获取：%v", user.Id, user.NickName, getValue, nowValue, addNum, newGetValue)
	//推送客户端消息
	this.GetUserManager().SendMessage(user, &pb.BagSpaceAddAck{BagMax: int32(user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum)}, true)
	return nil
}
