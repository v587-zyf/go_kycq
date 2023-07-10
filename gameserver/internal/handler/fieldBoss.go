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
	pb.Register(pb.CmdFieldBossLoadReqId, HandlerFieldBossLoadReq)
	pb.Register(pb.CmdFieldBossBuyNumReqId, HandlerFieldBossBuyNumReq)
	pb.Register(pb.CmdEnterFieldBossFightReqId, HandlerEnterFieldBossFightReq)
	pb.Register(pb.CmdFieldBossFirstReqId, HandlerFieldBossFirstReq)
}

func HandlerFieldBossLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FieldBossLoadReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.FieldBossLoadAck{}
	err = m.FieldBoss.Load(user, int(req.Area), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFieldBossLoadReq ack:%v", ack)
	return ack, nil, nil
}

func HandlerFieldBossBuyNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FieldBossBuyNumReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFieldBossFight)

	var err error
	ack := &pb.FieldBossBuyNumAck{}
	err = m.FieldBoss.BuyNum(user, req.Use, int(req.BuyNum), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFieldBossBuyNumReq ack:%v", ack)
	return ack, op, nil
}

func HandlerEnterFieldBossFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterFieldBossFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.FieldBoss.EnterFieldBossFight(user, int(req.StageId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.EnterFieldBossFightAck{StageId: int32(req.StageId), DareNum: int32(user.FieldBoss.DareNum)}, nil, nil
}

func HandlerFieldBossFirstReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFieldBossFirstReceive)
	if err := m.FieldBoss.FirstReceive(user, op); err != nil {
		return nil, nil, err
	}
	return &pb.FieldBossFirstAck{FirstReceive: true, Goods: op.ToChangeItems()}, op, nil
}
