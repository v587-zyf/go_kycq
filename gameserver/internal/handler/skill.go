package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"time"
)

func init() {
	pb.Register(pb.CmdSkillUpLvReqId, HandlerSkillUpLvReq)
	pb.Register(pb.CmdSkillChangePosReqId, HandlerSkillChangePosReq)
	pb.Register(pb.CmdSkillChangeWearReqId, HandlerSkillChangeWearReq)
	pb.Register(pb.CmdSkillResetReqId, HandlerSkillResetReq)
	pb.Register(pb.CmdSkillUseReqId, HandlerSkillUseReq)

	pb.Register(pb.CmdCutTreasureUpLvReqId, HandlerCutTreasureUpLvReq)
	pb.Register(pb.CmdCutTreasureUseReqId, HandlerCutTreasureUseReq)

	pb.Register(pb.CmdAncientSkillActiveReqId, HandlerAncientSkillActiveReq)
	pb.Register(pb.CmdAncientSkillUpLvReqId, HandlerAncientSkillUpLvReq)
	pb.Register(pb.CmdAncientSkillUpGradeReqId, HandlerAncientSkillUpGradeReq)
}

func HandlerSkillUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SkillUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeSKillUpLv)

	var err error
	ack := &pb.SkillUpLvAck{}

	err = m.Skill.UpLv(user, op, req, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerSkillUpLvReq ack is %v", ack)

	return ack, op, nil
}

func HandlerSkillChangePosReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SkillChangePosReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.SkillChangePosAck{}

	err = m.Skill.ChangePos(user, req, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerSkillChangePosReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerSkillChangeWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SkillChangeWearReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.SkillChangeWearAck{}

	err = m.Skill.ChangeWear(user, req, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerSkillChangeWearReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerSkillResetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SkillResetReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeSkillReset)

	var err error
	ack := &pb.SkillResetAck{}

	err = m.Skill.Reset(user, op, req, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerSkillReset ack is %v", ack)

	return ack, op, nil
}

func HandlerSkillUseReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SkillUseReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	err = m.Skill.UseSkill(user, req)
	if err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}

func HandlerCutTreasureUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCutTreasure)

	err := m.Skill.CutTreasureUpLv(user, op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.CutTreasureUpLvAck{CutTreasureLv: int32(user.CutTreasure)}, op, nil
}

func HandlerCutTreasureUseReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	endCd, err := m.Skill.CutTreasureSkillUse(user)
	if err != nil {
		return nil, nil, err
	}

	return &pb.CutTreasureUseAck{UseTime: int32(time.Now().Unix()), CdEndTime: int32(endCd)}, nil, nil
}

func HandlerAncientSkillActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientSkillActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientSkillActive)

	if err := m.Skill.AncientSkillActive(user, int(req.HeroIndex), op); err != nil {
		return nil, nil, err
	}
	return &pb.AncientSkillActiveAck{
		HeroIndex:    req.HeroIndex,
		AncientSkill: builder.BuildAncientSkill(user.Heros[int(req.HeroIndex)]),
	}, op, nil
}

func HandlerAncientSkillUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientSkillUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientSkillUpLv)

	if err := m.Skill.AncientSkillUpLv(user, int(req.HeroIndex), op); err != nil {
		return nil, nil, err
	}
	return &pb.AncientSkillUpLvAck{
		HeroIndex: req.HeroIndex,
		Level:     int32(user.Heros[int(req.HeroIndex)].AncientSkill.Level),
	}, op, nil
}

func HandlerAncientSkillUpGradeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AncientSkillUpGradeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAncientSkillUpGrade)

	if err := m.Skill.AncientSkillUpGrade(user, int(req.HeroIndex), op); err != nil {
		return nil, nil, err
	}
	return &pb.AncientSkillUpGradeAck{
		HeroIndex: req.HeroIndex,
		Grade:     int32(user.Heros[int(req.HeroIndex)].AncientSkill.Grade),
	}, op, nil
}
