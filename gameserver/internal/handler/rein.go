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
	pb.Register(pb.CmdReinActiveReqId, HandlerReinActiveReq)
	pb.Register(pb.CmdReincarnationReqId, HandlerReincarnationReq)
	pb.Register(pb.CmdReinCostBuyReqId, HandlerReinCostBuyReq)
	pb.Register(pb.CmdReinCostUseReqId, HandlerReinCostUseReq)
}

func HandlerReinActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.ReinActiveAck{}

	err = m.Rein.ReinActive(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerReinActiveReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerReincarnationReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.ReincarnationAck{}

	err = m.Rein.Reincarnation(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerReincarnationReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerReinCostBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ReinCostBuyReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeReinCostBuy)

	var err error
	ack := &pb.ReinCostBuyAck{}

	err = m.Rein.ReinCostBuy(user, int(req.Id), int(req.Num), req.Use, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerReinCostBuyReq ack is %v", ack)
	return ack, op, nil
}

func HandlerReinCostUseReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ReinCostUseReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeReinCostUse)

	var err error
	ack := &pb.ReinCostUseAck{}

	err = m.Rein.ReinCostUse(user, int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerReinCostUseReq ack is %v", ack)
	return ack, op, nil
}
