package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdWarOrderTaskFinishReqId, HandlerWarOrderTaskFninshReq)
	pb.Register(pb.CmdWarOrderTaskRewardReqId, HandlerWarOrderTaskRewardReq)
	//pb.Register(pb.CmdWarOrderBuyLuxuryReqId, HandlerWarOrderBuyLuxuryReq)
	//pb.Register(pb.CmdWarOrderBuyExpReqId, HandlerWarOrderBuyExpReq)
	pb.Register(pb.CmdWarOrderLvRewardReqId, HandlerWarOrderLvRewardReq)
	pb.Register(pb.CmdWarOrderExchangeReqId, HandlerWarOrderExchange)
	pb.Register(pb.CmdWarOrderOpenReqId, HandlerWarOrderOpenReq)
}

func HandlerWarOrderTaskFninshReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WarOrderTaskFinishReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWarOrderTaskFininsh)

	var err error
	ack := &pb.WarOrderTaskFinishAck{}

	err = m.WarOrder.TaskFinish(user, op, req, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerWarOrderTaskFninshReq ack is %v", ack)

	return ack, op, nil
}

func HandlerWarOrderTaskRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WarOrderTaskRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWarOrderTaskReward)

	var err error
	ack := &pb.WarOrderTaskRewardAck{}

	err = m.WarOrder.TaskReward(user, op, req, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerWarOrderTaskRewardReq ack is %v", ack)

	return ack, op, nil
}

//func HandlerWarOrderBuyLuxuryReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	user := conn.GetSession().(*managers.ClientSession).User
//	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWarOrderBuyLuxury)
//
//	var err error
//	ack := &pb.WarOrderBuyLuxuryAck{}
//
//	err = m.WarOrder.BuyLuxury(user, op, ack)
//	if err != nil {
//		return nil, nil, err
//	}
//	logger.Debug("HandlerWarOrderBuyLuxuryReq ack is %v", ack)
//
//	return ack, op, nil
//}

//func HandlerWarOrderBuyExpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	user := conn.GetSession().(*managers.ClientSession).User
//	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWarOrderBuyLuxury)
//
//	var err error
//	ack := &pb.WarOrderBuyExpAck{}
//
//	err = m.WarOrder.BuyExp(user, op, ack)
//	if err != nil {
//		return nil, nil, err
//	}
//	logger.Debug("HandlerWarOrderBuyExpReq ack is %v", ack)
//
//	return ack, op, nil
//}

func HandlerWarOrderLvRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WarOrderLvRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWarOrderLvReward)

	var err error
	ack := &pb.WarOrderLvRewardAck{}

	err = m.WarOrder.LvReward(user, int(req.Lv), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerWarOrderLvRewardReq ack is %v", ack)

	return ack, op, nil
}

func HandlerWarOrderExchange(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WarOrderExchangeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWarOrderExchange)

	var err error
	ack := &pb.WarOrderExchangeAck{}

	err = m.WarOrder.Exchange(user, int(req.ExchangeId), int(req.Num), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerWarOrderExchange ack is %v", ack)

	return ack, op, nil
}

func HandlerWarOrderOpenReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	userWarOrder := user.WarOrder
	return &pb.WarOrderOpenAck{WarOrderInfo: &pb.WarOrderTaskNtf{
		Task:     &pb.WarOrderTask{Task: builder.BuildWarOrderTask(userWarOrder.Task)},
		WeekTask: builder.BuildWarOrderWeekTask(userWarOrder.WeekTask),
		Lv:       int32(userWarOrder.Lv),
		Exp:      int32(userWarOrder.Exp),
	}}, nil, nil
}
