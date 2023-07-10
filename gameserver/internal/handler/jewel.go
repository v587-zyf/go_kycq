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
	pb.Register(pb.CmdJewelMakeReqId, HandlerJewelMakeReq)
	pb.Register(pb.CmdJewelUpLvReqId, HandlerJewelUpLvReq)
	pb.Register(pb.CmdJewelChangeReqId, HandlerJewelChangeReq)
	pb.Register(pb.CmdJewelRemoveReqId, HandlerJewelRemoveReq)
	pb.Register(pb.CmdJewelMakeAllReqId, HandlerJewelMakeAllReq)
}

func HandlerJewelMakeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.JewelMakeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeJewel)

	var err error
	ack := &pb.JewelMakeAck{}

	err = m.Jewel.JewelMake(user, int(req.HeroIndex), int(req.EquipPos), int(req.JewelPos), int(req.ItemId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerJewelMakeReq ack is %v", ack)

	return ack, op, nil
}

func HandlerJewelUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.JewelUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeJewel)

	var err error
	ack := &pb.JewelUpLvAck{}

	err = m.Jewel.JewelUpLv(user, int(req.HeroIndex), int(req.EquipPos), int(req.JewelPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerJewelUpLvReq ack is %v", ack)

	return ack, op, nil
}

func HandlerJewelChangeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.JewelChangeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeJewel)

	var err error
	ack := &pb.JewelChangeAck{}

	err = m.Jewel.JewelChange(user, int(req.HeroIndex), int(req.EquipPos), int(req.JewelPos), int(req.ItemId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerJewelChangeReq ack is %v", ack)

	return ack, op, nil
}

func HandlerJewelRemoveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.JewelRemoveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeJewel)

	var err error
	ack := &pb.JewelRemoveAck{}

	err = m.Jewel.JewelRemove(user, int(req.HeroIndex), int(req.EquipPos), int(req.JewelPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerJewelRemoveReq ack is %v", ack)

	return ack, op, nil
}

func HandlerJewelMakeAllReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.JewelMakeAllReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeJewel)

	var err error
	ack := &pb.JewelMakeAllAck{}

	err = m.Jewel.JewelMakeAll(user, int(req.HeroIndex), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerJewelMakeAllReq ack is %v", ack)

	return ack, op, nil
}
