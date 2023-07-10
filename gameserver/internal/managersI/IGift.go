package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"time"
)

type IGiftManager interface {
	Online(user *objs.User)

	OpenGift(user *objs.User, types, giftItemId, num int, chooseItemId []int, op *ophelper.OpBagHelperDefault) error

	//礼包码
	GiftCodeReward(user *objs.User, code string, op *ophelper.OpBagHelperDefault) error
	GiftCodeRewardLocal(user *objs.User, code string, op *ophelper.OpBagHelperDefault) error

	//限时礼包
	SendLimitedGift(user *objs.User) *pb.LimitedGiftNtf
	LimitedGiftSend(user *objs.User)
	LimitedGiftBuy(user *objs.User, moduleType int, op *ophelper.OpBagHelperDefault) error
	LimitedGiftCheckBuy(user *objs.User, typeId, payNum int) (string, error)
	LimitedGiftBuyOperation(user *objs.User, typeId int, op *ophelper.OpBagHelperDefault)
	GetNowLimitedGiftInfo(user *objs.User, typeId int) (*model.LimitGift, *model.LimitGiftUnit, gamedb.ItemInfos, int, int, int, string, error)

	//开服礼包
	OpenGiftBuyCheck(user *objs.User, giftId, money int) error
	OpenGiftBuyOperation(user *objs.User, giftId int, op *ophelper.OpBagHelperDefault)
	OpenGiftEndTime(user *objs.User) time.Time
}
