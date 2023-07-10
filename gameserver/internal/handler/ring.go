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
	pb.Register(pb.CmdRingWearReqId, HandlerRingWearReq)
	pb.Register(pb.CmdRingRemoveReqId, HandlerRingRemoveReq)
	pb.Register(pb.CmdRingStrengthenReqId, HandlerRingStrengthenReq)
	pb.Register(pb.CmdRingPhantomReqId, HandlerRingPhantomReq)
	pb.Register(pb.CmdRingSkillUpReqId, HandlerRingSkillUpReq)
	pb.Register(pb.CmdRingFuseReqId, HandlerRingFuseReq)
	pb.Register(pb.CmdRingSkillResetReqId, HandlerRingSkillResetReq)
}

func HandlerRingWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RingWearReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeRing)

	var err error
	ack := &pb.RingWearAck{}

	err = m.Ring.Wear(user, int(req.HeroIndex), int(req.RingPos), int(req.BagPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRingWearReq ack is %v", ack)
	return ack, op, nil
}

func HandlerRingRemoveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RingRemoveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeRing)

	var err error
	ack := &pb.RingRemoveAck{}

	err = m.Ring.Remove(user, int(req.HeroIndex), int(req.RingPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRingRemoveReq ack is %v", ack)
	return ack, op, nil
}

func HandlerRingStrengthenReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RingStrengthenReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeRing)

	var err error
	ack := &pb.RingStrengthenAck{}

	err = m.Ring.Strengthen(user, int(req.HeroIndex), int(req.RingPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRingStrengthenReq ack is %v", ack)
	return ack, op, nil
}

func HandlerRingPhantomReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RingPhantomReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeRing)

	var err error
	ack := &pb.RingPhantomAck{}

	err = m.Ring.RingPhantom(user, int(req.HeroIndex), int(req.RingPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRingPhantomReq ack is %v", ack)
	return ack, op, nil
}

func HandlerRingSkillUpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RingSkillUpReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.RingSkillUpAck{}

	err = m.Ring.PhantomSkill(user, int(req.HeroIndex), int(req.RingPos), int(req.PhantomPos), int(req.SkillId), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRingSkillUpReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerRingFuseReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RingFuseReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeRing)

	var err error
	ack := &pb.RingFuseAck{}

	err = m.Ring.Fuse(user, int(req.Id), int(req.BagPos1), int(req.BagPos2), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRingFuseReq ack is %v", ack)
	return ack, op, nil
}

func HandlerRingSkillResetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.RingSkillResetReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.RingSkillResetAck{}

	err = m.Ring.ResetSkill(user, int(req.HeroIndex), int(req.RingPos), int(req.PhantomPos), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRingSkillResetReq ack is %v", ack)
	return ack, nil, nil
}
