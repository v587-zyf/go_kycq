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
	pb.Register(pb.CmdTalentUpLvReqId, HandlerTalentUpLvReq)
	pb.Register(pb.CmdTalentResetReqId, HandlerTalentResetReq)
}

func HandlerTalentUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TalentUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.TalentUpLvAck{}
	err := m.Talent.UpLv(user, int(req.HeroIndex), int(req.Id), req.IsMax, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerTalentUpLvReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerTalentResetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.TalentResetReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTalent)

	ack := &pb.TalentResetAck{}
	err := m.Talent.Reset(user, int(req.HeroIndex), int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerTalentResetReq ack is %v", ack)

	return ack, op, nil
}
