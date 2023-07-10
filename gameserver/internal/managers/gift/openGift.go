package gift

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"fmt"
	"time"
)

/**
 *  @Description: 开服礼包购买校验
 *  @param user
 *  @param giftId	礼包id
 *  @param money	金额
 *  @return error
 */
func (this *GiftManager) OpenGiftBuyCheck(user *objs.User, giftId, money int) error {
	openGiftCfg := gamedb.GetOpenGiftOpenGiftCfg(giftId)
	if openGiftCfg == nil {
		return gamedb.ERRPARAM
	}
	userOpenGift := user.OpenGift
	buyNum, ok := userOpenGift[giftId]
	if !ok {
		buyNum = 0
	}
	if buyNum >= openGiftCfg.Time {
		return gamedb.ERRBUYTIMESLIMIT
	}
	if money != openGiftCfg.Price {
		return gamedb.ERRBUYNUM
	}
	if time.Now().Unix() > this.OpenGiftEndTime(user).Unix() {
		return gamedb.ERRACTIVITYCLOSE
	}
	return nil
}

/**
 *  @Description: 开服礼包购买后续操作
 *  @param user
 *  @param giftId	礼包id
 *  @param op
 */
func (this *GiftManager) OpenGiftBuyOperation(user *objs.User, giftId int, op *ophelper.OpBagHelperDefault) {
	openGiftCfg := gamedb.GetOpenGiftOpenGiftCfg(giftId)
	userOpenGift := user.OpenGift
	buyNum, ok := userOpenGift[giftId]
	if !ok {
		buyNum = 0
	}
	if buyNum >= openGiftCfg.Time {
		return
	}

	userOpenGift[giftId]++
	user.Dirty = true
	this.GetBag().AddItems(user, openGiftCfg.Item, op)

	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_BUY_OPEN_GIFT, []int{1})
	this.GetAnnouncement().SendSystemChat(user, pb.SCROLINGTYPE_KAI_FU_BUY, -1, -1)
	this.GetUserManager().SendMessage(user, &pb.OpenGiftBuyNtf{Goods: op.ToChangeItems(), OpenGiftId: int32(giftId), BuyNum: int32(userOpenGift[giftId])}, true)
}

/**
 *  @Description: 获取开服礼包结束时间
 *  @param user
 *  @return time.Time
 */
func (this *GiftManager) OpenGiftEndTime(user *objs.User) time.Time {
	serverOpenTime := this.GetSystem().GetServerOpenTimeByServerId(user.ServerId)
	openGiftTime := gamedb.GetConf().OpenGiftTime
	addHour, _ := time.ParseDuration(fmt.Sprintf(`%dh`, openGiftTime[1]))
	addMinute, _ := time.ParseDuration(fmt.Sprintf(`%dm`, openGiftTime[2]))
	endTime := serverOpenTime.AddDate(0, 0, openGiftTime[0]).Add(addHour).Add(addMinute)
	return endTime
}
