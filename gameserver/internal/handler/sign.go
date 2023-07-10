package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdSignReqId, HandlerSignReq)
	pb.Register(pb.CmdSignRepairReqId, HandlerSignRepairReq)
	pb.Register(pb.CmdCumulativeSignReqId, HandlerCumulativeSignReq)
}

func HandlerSignReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeSign)

	ack := &pb.SignAck{}
	err := m.Sign.Sign(user, op, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerSignRepairReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SignRepairReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeSign)

	ack := &pb.SignRepairAck{}
	err := m.Sign.Repair(user, int(req.RepairDay), op, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerCumulativeSignReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.CumulativeSignReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeSign)

	ack := &pb.CumulativeSignAck{}
	err := m.Sign.Cumulative(user, int(req.CumulativeDay), op, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}
