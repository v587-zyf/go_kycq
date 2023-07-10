package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdWingUpLevelReqId, HandlerWingUpLevelReq)
	pb.Register(pb.CmdWingSpecialUpReqId, HandlerWingSpecialUpReq)
	pb.Register(pb.CmdWingUseMaterialReqId, HandlerWingUseMaterialReq)
	pb.Register(pb.CmdWingWearReqId, HandlerWingWearReq)
}

func HandlerWingUpLevelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WingUpLevelReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWingUpLevel)

	var err error
	ack := &pb.WingUpLevelAck{}

	err = m.Wing.UpLevel(user, int(req.HeroIndex), op, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerWingSpecialUpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WingSpecialUpReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWingSpecialUp)

	var err error
	ack := &pb.WingSpecialUpAck{}

	err = m.Wing.UpSpecialLevel(user, int(req.HeroIndex), int(req.SpecialType), op, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerWingUseMaterialReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WingUseMaterialReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWingUpLevel)

	var err error
	ack := &pb.WingUseMaterialAck{}

	err = m.Wing.UseMaterial(user, int(req.HeroIndex), op, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerWingWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WingWearReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.WingWearAck{}

	err = m.Wing.Wear(user, int(req.HeroIndex), int(req.WingId), ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, nil, nil
}
