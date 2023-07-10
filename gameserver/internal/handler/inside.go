package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdInsideUpStarReqId, HandlerInsideUpStarReq)
	pb.Register(pb.CmdInsideUpGradeReqId, HandlerInsideUpGradeReq)
	pb.Register(pb.CmdInsideUpOrderReqId, HandlerInsideUpOrderReq)
	pb.Register(pb.CmdInsideSkillUpLvReqId, HandlerInsideSkillUpLvReq)
	pb.Register(pb.CmdInsideAutoUpReqId, HandlerInsideAutoUpReq)
}

func HandlerInsideUpStarReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.InsideUpStarReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeInside)

	var err error
	err = m.Inside.InsideUpStar(user, int(req.HeroIndex), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.InsideUpStarAck{HeroIndex: req.HeroIndex, InsideInfo: builder.BuildInsideInfo(user.Heros[int(req.HeroIndex)])}, op, nil
}

func HandlerInsideUpGradeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.InsideUpGradeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeInside)

	var err error
	ack := &pb.InsideUpGradeAck{}

	err = m.Inside.InsideUpGrade(user, int(req.HeroIndex), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerInsideUpGradeReq ack is %v", ack)

	return ack, op, nil
}

func HandlerInsideUpOrderReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.InsideUpOrderReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeInside)

	var err error
	ack := &pb.InsideUpOrderAck{}

	err = m.Inside.InsideUpOrder(user, int(req.HeroIndex), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerInsideUpOrderReq ack is %v", ack)

	return ack, op, nil
}

func HandlerInsideSkillUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.InsideSkillUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeInside)

	var err error
	ack := &pb.InsideSkillUpLvAck{}

	err = m.Inside.InsideSkillUpLv(user, int(req.HeroIndex), int(req.SkillId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerInsideSkillUpLvReq ack is %v", ack)

	return ack, op, nil
}

func HandlerInsideAutoUpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.InsideAutoUpReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeInsideAutoUp)

	var err error
	err = m.Inside.AutoUp(user, int(req.HeroIndex), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.InsideAutoUpAck{HeroIndex: req.HeroIndex, InsideInfo: builder.BuildInsideInfo(user.Heros[int(req.HeroIndex)])}, op, nil
}
