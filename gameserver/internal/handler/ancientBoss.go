package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdAncientBossLoadReqId, HandlerAncientBossLoadReq)
	pb.Register(pb.CmdAncientBossBuyNumReqId, HandlerAncientBossBuyNumReq)
	pb.Register(pb.CmdEnterAncientBossFightReqId, HandlerEnterAncientBossFightReq)
	pb.Register(pb.CmdAncientBossOwnerReqId, HandlerAncientBossOwnerReq)
}

func HandlerAncientBossLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientBossLoadReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.AncientBossLoadAck{}
	err := m.AncientBoss.Load(user, int(req.Area), ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, nil, nil
}

func HandlerAncientBossBuyNumReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientBossBuyNumReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientBuyNum)

	err := m.AncientBoss.BuyNum(user, req.Use, int(req.BuyNum), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.AncientBossBuyNumAck{
		BuyNum: int32(user.AncientBoss.BuyNum),
	}, op, nil
}

func HandlerEnterAncientBossFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterAncientBossFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.AncientBoss.EnterAncientBossFight(user, int(req.StageId))
	if err != nil {
		return nil, nil, err
	}

	return &pb.EnterAncientBossFightAck{
		StageId: req.StageId,
		DareNum: int32(user.AncientBoss.DareNum),
	}, nil, nil
}

func HandlerAncientBossOwnerReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientBossOwnerReq)
	user := conn.GetSession().(*managers.ClientSession).User

	owner := m.AncientBoss.GetOwnerList(user, int(req.StageId))
	return &pb.AncientBossOwnerAck{
		StageId: req.StageId,
		List:    owner,
	}, nil, nil
}
