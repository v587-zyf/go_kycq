package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdFitUpLvReqId, HandlerFitUpLvReq)
	pb.Register(pb.CmdFitSkillActiveReqId, HandlerFitSkillActiveReq)
	pb.Register(pb.CmdFitSkillUpLvReqId, HandlerFitSkillUpLvReq)
	pb.Register(pb.CmdFitSkillUpStarReqId, HandlerFitSkillUpStarReq)
	pb.Register(pb.CmdFitSkillChangeReqId, HandlerFitSkillChangeReq)
	pb.Register(pb.CmdFitSkillResetReqId, HandlerFitSkillResetReq)
	pb.Register(pb.CmdFitFashionUpLvReqId, HandlerFitFashionUpLvReq)
	pb.Register(pb.CmdFitFashionChangeReqId, HandlerFitFashionChangeReq)
	pb.Register(pb.CmdFitEnterReqId, HandlerFitEnter)
	pb.Register(pb.CmdFitCancleReqId, HandlerFitCancel)

	pb.Register(pb.CmdFitHolyEquipComposeReqId, HandlerFitHolyEquipComposeReq)
	pb.Register(pb.CmdFitHolyEquipDeComposeReqId, HandlerFitHolyEquipDeComposeReq)
	pb.Register(pb.CmdFitHolyEquipWearReqId, HandlerFitHolyEquipWearReq)
	pb.Register(pb.CmdFitHolyEquipRemoveReqId, HandlerFitHolyEquipRemoveReq)
	pb.Register(pb.CmdFitHolyEquipSuitSkillChangeReqId, HandlerFitHolyEquipSuitSkillChangeReq)
}

func HandlerFitHolyEquipComposeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitHolyEquipComposeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFitHolyEquipCompose)

	var err error
	ack := &pb.FitHolyEquipComposeAck{}

	err = m.Fit.HolyEquipCompose(user, int(req.EquipId), int(req.Pos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitHolyEquipComposeReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitHolyEquipDeComposeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitHolyEquipDeComposeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFitHolyEquipDeCompose)

	var err error
	ack := &pb.FitHolyEquipDeComposeAck{}

	err = m.Fit.HolyEquipDeCompose(user, int(req.BagPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitHolyEquipDeComposeReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitHolyEquipWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitHolyEquipWearReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFitHolyEquipWear)

	var err error
	ack := &pb.FitHolyEquipWearAck{}

	err = m.Fit.HolyEquipWear(user, int(req.BagPos), int(req.EquipPos), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitHolyEquipWearReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitHolyEquipRemoveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitHolyEquipRemoveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFitHolyEquipWear)

	var err error
	ack := &pb.FitHolyEquipRemoveAck{}

	err = m.Fit.HolyEquipRemove(user, int(req.Pos), int(req.SuitType), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitHolyEquipRemoveReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitHolyEquipSuitSkillChangeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitHolyEquipSuitSkillChangeReq)
	user := conn.GetSession().(*managers.ClientSession).User

	err := m.Fit.HolyEquipSuitSkillChange(user, int(req.SuitId))
	if err != nil {
		return nil, nil, err
	}
	return &pb.FitHolyEquipSuitSkillChangeAck{SuitId: req.SuitId}, nil, nil
}

func HandlerFitUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFit)

	var err error
	ack := &pb.FitUpLvAck{}

	err = m.Fit.UpLv(user, int(req.FitId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitUpLvReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitSkillActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitSkillActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFit)

	var err error
	ack := &pb.FitSkillActiveAck{}

	err = m.Fit.SkillActive(user, int(req.FitSkillId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitSkillActiveReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitSkillUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitSkillUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFit)

	var err error
	ack := &pb.FitSkillUpLvAck{}

	err = m.Fit.SkillUpLv(user, int(req.FitSkillId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitSkillUpLvReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitSkillUpStarReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitSkillUpStarReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFit)

	var err error
	ack := &pb.FitSkillUpStarAck{}

	err = m.Fit.SkillUpStar(user, int(req.FitSkillId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitSkillUpStarReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitSkillChangeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitSkillChangeReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.FitSkillChangeAck{}

	err = m.Fit.SkillChange(user, int(req.FitSkillId), int(req.FitSkillSlot), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitSkillChangeReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerFitSkillResetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitSkillResetReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFit)

	var err error
	ack := &pb.FitSkillResetAck{}

	err = m.Fit.SkillReset(user, int(req.FitSkillId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitSkillResetReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitFashionUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitFashionUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFit)

	var err error
	ack := &pb.FitFashionUpLvAck{}

	err = m.Fit.FashionUpLv(user, int(req.FitFashionId), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitFashionUpLvReq ack is %v", ack)
	return ack, op, nil
}

func HandlerFitFashionChangeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.FitFashionChangeReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.FitFashionChangeAck{}

	err = m.Fit.FashionChange(user, int(req.FitFashionId), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerFitFashionChangeReq ack is %v", ack)
	return ack, nil, nil
}

func HandlerFitEnter(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.FitEnterAck{}
	err := m.Fit.EnterFit(user, constFight.FIT_ID, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

func HandlerFitCancel(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Fit.FitCancel(user)
	if err != nil {
		return nil, nil, err
	}
	return &pb.FitCancleAck{}, nil, nil
}
