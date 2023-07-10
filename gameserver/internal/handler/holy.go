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
	pb.Register(pb.CmdHolyActiveReqId, HandlerHolyActiveReq)
	pb.Register(pb.CmdHolyUpLevelReqId, HandlerHolyUpLevelReq)
	pb.Register(pb.CmdHolySkillActiveReqId, HandlerHolySkillActiveReq)
	pb.Register(pb.CmdHolySkillUpLvReqId, HandlerHolySkillUpLvReq)
}

func HandlerHolyActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.HolyActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.HolyActiveAck{}

	if err = m.Holyarms.Active(user, int(req.Id), ack); err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerHolyActiveReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerHolyUpLevelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.HolyUpLevelReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHolyarms)

	var err error
	ack := &pb.HolyUpLevelAck{}

	if err = m.Holyarms.AutoUpLv(user, int(req.Id), op, ack); err != nil {
	//if err = m.Holyarms.UpLevel(user, int(req.Id), int(req.ItemId), op, ack); err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerHolyUpLevelReq ack is %v", ack)
	return ack, op, nil
}

func HandlerHolySkillActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.HolySkillActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHolyarms)

	var err error
	ack := &pb.HolySkillActiveAck{}

	if err = m.Holyarms.ActiveSkill(user, int(req.Hid), int(req.Hlv), op, ack); err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerHolySkillActiveReq ack is %v", ack)
	return ack, op, nil
}

func HandlerHolySkillUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.HolySkillUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeHolyarms)

	var err error
	ack := &pb.HolySkillUpLvAck{}

	if err = m.Holyarms.SkillUpLv(user, int(req.Hid), int(req.Hlv), op, ack); err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerHolySkillUpLvReq ack is %v", ack)
	return ack, op, nil
}
