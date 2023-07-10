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
	pb.Register(pb.CmdVipBossLoadReqId, HandlerVipBossLoadReq)
	pb.Register(pb.CmdEnterVipBossFightReqId, HandlerEnterVipBossFightReq)
	pb.Register(pb.CmdVipBossSweepReqId, HandlerVipBossSweepReq)
}

func HandlerVipBossLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	return &pb.VipBossLoadAck{VipBoss: builder.BuildVipBoss(user)}, nil, nil
}

func HandlerEnterVipBossFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterVipBossFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.VipBoss.EnterVipBossFight(user, int(req.StageId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.EnterVipBossFightAck{StageId: int32(req.StageId), DareNum: int32(user.VipBosses.DareNum[int(req.StageId)])}, nil, nil
}

func HandlerVipBossSweepReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.VipBossSweepReq)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeVipBossSweep)

	ack := &pb.VipBossSweepAck{}
	err := m.VipBoss.VipBossSweep(user, int(req.StageId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
