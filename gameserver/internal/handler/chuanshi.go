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
	pb.Register(pb.CmdChuanShiWearReqId, HandlerChuanShiWearReq)
	pb.Register(pb.CmdChuanShiRemoveReqId, HandlerChuanshiRemoveReq)
	pb.Register(pb.CmdChuanShiDeComposeReqId, HandlerChuanshiDeComposeReq)

	pb.Register(pb.CmdChuanshiStrengthenReqId, HandlerChuanShiStrengthenReq)
}

func HandlerChuanShiWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ChuanShiWearReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeChuanShiEquipWear)

	ack := &pb.ChuanShiWearAck{}
	err := m.ChuanShi.Wear(user, int(req.HeroIndex), int(req.BagPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerChuanShiWearReq ack is %v", ack)

	m.UserManager.SendItemChangeNtf(user, op)
	return ack, nil, nil
}

func HandlerChuanshiRemoveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ChuanShiRemoveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeChuanShiEquipRemove)

	err := m.ChuanShi.Remove(user, int(req.HeroIndex), int(req.EquipPos), op)
	if err != nil {
		return nil, nil, err
	}

	m.UserManager.SendItemChangeNtf(user, op)
	return &pb.ChuanShiRemoveAck{HeroIndex: req.HeroIndex, EquipPos: req.EquipPos}, nil, nil
}

func HandlerChuanshiDeComposeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ChuanShiDeComposeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeChuanShiEquipDeCompose)

	err := m.ChuanShi.DeCompose(user, int(req.BagPos), op)
	if err != nil {
		return nil, nil, err
	}
	return &pb.ChuanShiDeComposeAck{Goods: op.ToChangeItems()}, op, nil
}

func HandlerChuanShiStrengthenReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ChuanshiStrengthenReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeChuanShiStrengthen)

	ack := &pb.ChuanshiStrengthenAck{}
	err := m.ChuanShi.Strengthen(user, int(req.HeroIndex), int(req.EquipPos), int(req.Stone), op, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}
