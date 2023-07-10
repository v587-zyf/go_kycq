package gift

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"time"
)

/**
 *  @Description: 礼包码核销
 *  @param user
 *  @param code	礼包码
 *  @param op
 *  @return error
 */
func (this *GiftManager) GiftCodeReward(user *objs.User, code string, op *ophelper.OpBagHelperDefault) error {

	rewards, err := ptsdk.GetSdk().ExchangeCode(code, user.Id, user.NickName, user.ServerId,user.ChannelId)
	if err != nil {
		return err
	}
	items := make(gamedb.ItemInfos, len(rewards))
	index := 0
	for k, v := range rewards {
		items[index] = &gamedb.ItemInfo{ItemId: k, Count: v}
		index++
	}
	this.GetBag().AddItems(user, items, op)
	return nil
}

/**
 *  @Description: 礼包码核销（暂时不用）
 *  @param user
 *  @param code	礼包码
 *  @param op
 *  @return error
 */
func (this *GiftManager) GiftCodeRewardLocal(user *objs.User, code string, op *ophelper.OpBagHelperDefault) error {
	timeNow := time.Now()
	giftCodeModel := modelCross.GetGiftCodeDbModel()
	giftCodeReceiveModel := modelCross.GetGiftCodeReceiveDbModel()
	giftCodeData, err := giftCodeModel.GetGiftCodeByCode(code)
	if err != nil {
		return gamedb.ERRGIFTCODE
	}
	if giftCodeData.ServerId != 0 && giftCodeData.ServerId != base.Conf.ServerId {
		return gamedb.ERRGIFTCODE
	}
	//todo 判断渠道
	if giftCodeData.StartTime.Unix() > timeNow.Unix() || giftCodeData.EndTime.Unix() < timeNow.Unix() {
		return gamedb.ERRGIFTCODE
	}

	//每个账号只能领取一次这个礼包码
	rNum, _ := giftCodeReceiveModel.GetGiftCodeReceiveNumByCodeIdAndUserId(giftCodeData.Id, user.Id)
	if rNum > 0 {
		return gamedb.ERRREPEATRECEIVE
	}
	//每个账号批次领取数量校验
	receiveTimes, err := giftCodeReceiveModel.GetUserReceiveTimesByBatch(user.Id, giftCodeData.BatchId)
	if err != nil {
		logger.Warn("GiftCodeReward GetUserReceiveTimesByBatch err:%v", err)
	}
	if receiveTimes >= giftCodeData.BatchNum {
		return gamedb.ERRRECEIVEMAX
	}

	addItems := make(gamedb.ItemInfos, 0)
	for itemId, count := range giftCodeData.Reward {
		addItems = append(addItems, &gamedb.ItemInfo{
			ItemId: itemId,
			Count:  count,
		})
	}
	this.GetBag().AddItems(user, addItems, op)
	err = giftCodeReceiveModel.Create(&modelCross.GiftCodeReceiveDb{
		CodeId:      giftCodeData.Id,
		UserId:      user.Id,
		BatchId:     giftCodeData.BatchId,
		ReceiveTime: time.Now(),
	})
	if err != nil {
		logger.Warn("GiftCodeReward GiftCodeReceiveDb create userId:%v err:%v", user.Id, err)
	}
	user.Dirty = true
	return nil
}
