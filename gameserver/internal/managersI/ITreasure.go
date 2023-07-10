package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ITreasureManager interface {
	//设置转盘玩家上线弹框提示
	SetPopUp(user *objs.User, state int, ack *pb.SetTreasurePopUpStateAck)

	Reset(user *objs.User, isEnter bool)

	Load(user *objs.User, ack *pb.TreasureInfosAck) error

	//花费元宝购买道具
	BuyTreasureItem(user *objs.User, ack *pb.BuyTreasureItemAck, op *ophelper.OpBagHelperDefault) error

	//转盘奖励选择
	ChooseTreasureItem(user *objs.User, types int, indexReq []int32, isReplace, replaceIndex int, ack *pb.ChooseTreasureAwardAck) error

	//获取阶段奖励
	GetTreasureIntegralAward(user *objs.User, index int, ack *pb.GetTreasureIntegralAwardAck, op *ophelper.OpBagHelperDefault) error

	//开始抽奖
	ApplyGet(user *objs.User, ack *pb.TreasureApplyGetAck, op *ophelper.OpBagHelperDefault) error

	//充值检查
	PayCheck(user *objs.User, payNum int) error

	//充值回调
	PayCallBack(user *objs.User, payNum int, op *ophelper.OpBagHelperDefault)

	DrawLoad(user *objs.User, ack *pb.TreasureDrawInfoAck)

	SendTreasureMail()
}
