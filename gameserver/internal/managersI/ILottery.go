package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ILotteryManager interface {
	Load(user *objs.User, ack *pb.LotteryInfoAck)

	//接好运
	GetGoodLucky(user *objs.User, ack *pb.GetGoodLuckAck, op *ophelper.OpBagHelperDefault) error

	//购买份额
	LotteryBuyNums(user *objs.User, nums int, ack *pb.LotteryBuyNumsAck, op *ophelper.OpBagHelperDefault) error

	//结算
	OpenAward()

	//设置弹框状态
	SetPopState(user *objs.User, ack *pb.SetLotteryPopUpStateAck)

	//上线检查
	OnlineCheck(user *objs.User)

	//重置
	Reset()

	UserReset(user *objs.User)

	//未主动领取奖励的邮件发送
	SendReward()

	//获奖信息
	GetAwardInfo(user *objs.User, ack *pb.LotteryInfo1Ack)

	//领取投注奖励
	GetLotteryAward(user *objs.User, ack *pb.LotteryGetEndAwardAck, op *ophelper.OpBagHelperDefault) error
}
