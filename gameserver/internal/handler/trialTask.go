package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdTrialTaskInfoReqId, HandlerTrialTaskInfoReq)
	pb.Register(pb.CmdTrialTaskGetAwardReqId, HandlerTrialTaskGetAwardReq)
	pb.Register(pb.CmdTrialTaskGetStageAwardReqId, HandlerTrialTaskGetStageAwardReq)
}

func HandlerTrialTaskInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.TrialTaskInfoAck{}
	m.GetTrialTask().TrialTaskLoad(user, ack)
	return ack, nil, nil
}

func HandlerTrialTaskGetAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TrialTaskGetAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGetTrialTaskAward)

	ack := &pb.TrialTaskGetAwardAck{}
	if err := m.GetTrialTask().GetTrialTaskAward(user, int(req.Id), ack, op); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerTrialTaskGetStageAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TrialTaskGetStageAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGetTrialTaskStageAward)

	ack := &pb.TrialTaskGetStageAwardAck{}
	if err := m.GetTrialTask().GetStageAward(user, int(req.Id), ack, op); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
