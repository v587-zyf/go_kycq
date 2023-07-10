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
	pb.Register(pb.CmdChallengeInfoReqId, HandlerChallengeInfoReq)
	pb.Register(pb.CmdApplyChallengeReqId, HandlerApplyChallengeReq)                     //报名
	pb.Register(pb.CmdChallengeEachRoundPeopleReqId, HandlerChallengeEachRoundPeopleReq) //当前轮参赛的人数
	pb.Register(pb.CmdBottomPourReqId, HandlerBottomPourReq)                             //下注
}

func HandlerChallengeInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.ChallengeInfoAck{}
	err := m.Challenge.LoadInfo(user.ServerId, ack)
	if err != nil {
		logger.Error("HandlerChallengeInfoReq SetApplyUserInfo err:%v", err)
		return nil, nil, err
	}
	return ack, nil, nil
}

func HandlerApplyChallengeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.ApplyChallengeAck{}
	err := m.Challenge.SetApplyUserInfo(user, ack)
	if err != nil {
		logger.Error("HandlerApplyChallengeReq SetApplyUserInfo err:%v", err)
		return nil, nil, err
	}
	return ack, nil, nil
}

func HandlerChallengeEachRoundPeopleReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.ChallengeEachRoundPeopleAck{}
	err := m.Challenge.EachRoundPeople(user, ack)
	if err != nil {
		logger.Error("HandlerChallengeEachRoundPeopleReq SetApplyUserInfo err:%v", err)
		return nil, nil, err
	}
	return ack, nil, nil
}

func HandlerBottomPourReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.BottomPourReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeChallengeBottom)
	ack := &pb.BottomPourAck{}
	err := m.Challenge.BottomPour(user, int(req.UserId), op, ack)
	if err != nil {
		logger.Error("HandlerBottomPourReq err:%v", err)
		return nil, nil, err
	}
	return ack, op, nil
}
