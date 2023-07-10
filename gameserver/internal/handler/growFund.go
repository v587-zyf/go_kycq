package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	//pb.Register(pb.CmdGrowFundBuyReqId, HandlerGrowFundBuyReq)
	pb.Register(pb.CmdGrowFundRewardReqId, HandlerGrowFundRewardReq)
}

//func HandlerGrowFundBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	user := conn.GetSession().(*managers.ClientSession).User
//	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGrowFund)
//
//	var err error
//	ack := &pb.GrowFundBuyAck{}
//
//	err = m.GrowFund.Buy(user, op, ack)
//	if err != nil {
//		return nil, nil, err
//	}
//	logger.Debug("HandlerGrowFundBuyReq ack is %v", ack)
//	return ack, op, nil
//}

func HandlerGrowFundRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GrowFundRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGrowFund)

	var err error
	ack := &pb.GrowFundRewardAck{}

	err = m.GrowFund.Reward(user, int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerGrowFundRewardReq ack is %v", ack)
	return ack, op, nil
}
