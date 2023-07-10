package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IDailyRankManager interface {
	LoadRankReq(user *objs.User, ack *pb.DailyRankLoadAck)

	//发送每日排行奖励
	SendEndMail() error

	//记录每天战力增长节点限制
	ResAddState()

	GetAddState() bool

	//获取积分奖励
	GetMarkReward(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.DailyRankGetMarkRewardAck) error

	//购买每日排行礼包
	BuyDailyRankGift(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.DailyRankBuyGiftAck) error

	PayCheck(user *objs.User, payNum, typeId int) error

	PayCallBack(user *objs.User, payNum, typeId int, op *ophelper.OpBagHelperDefault)

	OnlineCheck(user *objs.User)
}
