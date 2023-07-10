package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdLotteryInfoReqId, HandlerLotteryInfoReq)
	pb.Register(pb.CmdLotteryBuyNumsReqId, HandlerLotteryBuyNumsReq)             //购买份额
	pb.Register(pb.CmdGetGoodLuckReqId, HandlerGetGoodLuckReq)                   //接好运
	pb.Register(pb.CmdSetLotteryPopUpStateReqId, HandlerSetLotteryPopUpStateReq) //弹窗状态
	pb.Register(pb.CmdLotteryInfo1ReqId, HandlerLotteryGetAwardInfoReq)          //领取奖励界面信息
	pb.Register(pb.CmdLotteryGetEndAwardReqId, HandlerLotteryGetEndAwardReq)     //领取投注奖励
}

func HandlerLotteryInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.LotteryInfoAck{}
	m.Lottery.Load(user, ack)
	return ack, nil, nil
}

func HandlerLotteryBuyNumsReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.LotteryBuyNumsReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeLotteryBuyNum)
	ack := &pb.LotteryBuyNumsAck{}
	err := m.Lottery.LotteryBuyNums(user, int(req.Num), ack, op)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerGetGoodLuckReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGetGoodLuckNum)
	ack := &pb.GetGoodLuckAck{}
	err := m.Lottery.GetGoodLucky(user, ack, op)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerSetLotteryPopUpStateReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	user.LotteryInfo.PopUpState = 1
	ack := &pb.SetLotteryPopUpStateAck{
		State: int32(1),
	}
	return ack, nil, nil
}

func HandlerLotteryGetAwardInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.LotteryInfo1Ack{}
	m.Lottery.GetAwardInfo(user, ack)
	return ack, nil, nil
}

func HandlerLotteryGetEndAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGetGetEndAward)
	ack := &pb.LotteryGetEndAwardAck{}
	err := m.Lottery.GetLotteryAward(user, ack, op)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
