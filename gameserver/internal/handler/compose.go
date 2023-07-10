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
	pb.Register(pb.CmdComposeReqId, HandlerComposeReq)
	pb.Register(pb.CmdComposeEquipReqId, HandlerComposeEquipReq)
	pb.Register(pb.CmdComposeChuanShiEquipReqId, HandlerComposeChuanShiEquipReq)
}

func HandlerComposeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ComposeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCompose)

	var err error
	ack := &pb.ComposeAck{}

	err = m.Compose.Compose(user, int(req.HeroIndex), int(req.SubId), int(req.ComposeNum), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerComposeReq ack is %v", ack)

	return ack, op, nil
}

func HandlerComposeEquipReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ComposeEquipReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeComposeEquip)

	var err error
	err = m.Compose.ComposeEquip(user, req, op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.ComposeEquipAck{
		ComposeEquipSubId: req.ComposeEquipSubId,
		IsLuckyStone:      req.IsLuckyStone,
		BigLuckyStone:     req.BigLuckyStone,
		Goods:             op.ToChangeItems(),
	}, op, nil
}

func HandlerComposeChuanShiEquipReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ComposeChuanShiEquipReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeComposeChuanShiEquip)

	err := m.Compose.ComposeChuanShiEquip(user, int(req.ComposeSubId), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.ComposeChuanShiEquipAck{
		ComposeSubId: req.ComposeSubId,
		Goods:        op.ToChangeItems(),
	}, op, nil
}
