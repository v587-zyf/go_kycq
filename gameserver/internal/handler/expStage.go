package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdExpStageFightReqId, HandlerExpStageFightReq)
	pb.Register(pb.CmdExpStageDoubleReqId, HandlerExpStageDoubleReq)
	pb.Register(pb.CmdExpStageSweepReqId, HandlerExpStageSweepReq)
	pb.Register(pb.CmdExpStageBuyNumReqId, HandlerExpStageBuyNumReq)
}

func HandlerExpStageFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ExpStageFightReq)
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.ExpStage.EnterExpStageFight(user, int(req.StageId))
	if err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}

func HandlerExpStageDoubleReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ExpStageDoubleReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeExpStage)

	ack := &pb.ExpStageDoubleAck{}
	err := m.ExpStage.Double(user, op, int(req.StageId), ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerExpStageSweepReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ExpStageSweepReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeExpStageSweep)

	ack := &pb.ExpStageSweepAck{}
	err := m.ExpStage.ExpStageSweep(user, int(req.StageId), op, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerExpStageBuyNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ExpStageBuyNumReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeExpStageBuyNum)

	err := m.ExpStage.ExpStageBuyNum(user, req.Use, op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.ExpStageBuyNumAck{BuyNum: int32(user.ExpStage.BuyNum)}, op, nil
}
