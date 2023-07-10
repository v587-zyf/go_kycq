package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdPersonBossLoadReqId, HandlerPersonBossLoadReq)
	pb.Register(pb.CmdEnterPersonBossFightReqId, HandlerEnterPersonBossFightReq)
	pb.Register(pb.CmdPersonBossSweepReqId, HandlerPersonBossSweepReq)
}

func HandlerPersonBossLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	return &pb.PersonBossLoadAck{PersonBoss: builder.BuildPersonBoss(user)}, nil, nil
}

func HandlerEnterPersonBossFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.EnterPersonBossFightReq)

	ack := &pb.EnterPersonBossFightAck{}
	err := m.PersonBoss.EnterPersonBossFightReq(user, int(req.StageId), ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

func HandlerPersonBossSweepReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.PersonBossSweepReq)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePersonBossSweep)

	ack := &pb.PersonBossSweepAck{}
	err := m.PersonBoss.PersonBossSweep(user, int(req.StageId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
