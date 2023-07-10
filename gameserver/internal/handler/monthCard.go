package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	//pb.Register(pb.CmdMonthCardBuyReqId, HandlerMonthCardBuyReq)
	pb.Register(pb.CmdMonthCardDailyRewardReqId, HandlerMonthCardDailyRewardReq)
}

//func HandlerMonthCardBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	req := p.(*pb.MonthCardBuyReq)
//	user := conn.GetSession().(*managers.ClientSession).User
//	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMonthCard)
//
//	ack := &pb.MonthCardBuyAck{}
//	err := m.MonthCard.Buy(user, int(req.Id), op, ack)
//	if err != nil {
//		return nil, nil, err
//	}
//	logger.Debug("HandlerMonthCardBuyReq ack:%v", ack)
//
//	return ack, op, nil
//}

func HandlerMonthCardDailyRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.MonthCardDailyRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMonthCard)
	err := m.MonthCard.DailyReward(user, int(req.MonthCardType), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.MonthCardDailyRewardAck{
		MonthCardType: req.MonthCardType,
		Goods:         op.ToChangeItems(),
	}, op, nil
}
