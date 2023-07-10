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
	pb.Register(pb.CmdAchievementLoadReqId, HandlerAchievementLoadReq)
	pb.Register(pb.CmdAchievementGetAwardReqId, HandlerAchievementGetAwardReq)
	pb.Register(pb.CmdActiveMedalReqId, HandlerActiveMedalReq)
}

func HandlerAchievementLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.AchievementLoadAck{}

	err = m.Achievement.Load(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerAchievementLoadReq ack is %v", ack)

	return ack, nil, nil
}

func HandlerAchievementGetAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AchievementGetAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeAchieveGetReward)

	var err error
	ack := &pb.AchievementGetAwardAck{}

	err = m.Achievement.GetAward(user, req.Id, ack, op)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerAchievementGetAwardReq ack is %v", ack)

	return ack, op, nil
}

func HandlerActiveMedalReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ActiveMedalReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.ActiveMedalAck{}

	err = m.Achievement.ActiveMedal(user, int(req.Id), ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerActiveMedalReq ack is %v", ack)

	return ack, nil, nil
}
