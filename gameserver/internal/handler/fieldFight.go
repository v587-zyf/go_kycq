package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdFieldFightLoadReqId, HandlerFieldFightLoadReq)
	pb.Register(pb.CmdBuyFieldFightChallengeTimesReqId, HandlerBuyCompetitveChallengeTimesReq)
	pb.Register(pb.CmdRefFieldFightRivalUserReqId, HandlerRefFieldFightRivalUserReq)
	pb.Register(pb.CmdBuyFieldFightChallengeTimesReqId, HandlerBuyFieldFightChallengeTimesReq)
	pb.Register(pb.CmdEnterFieldFightReqId, HandlerEnterFieldFightReq)
}

//打开页面
func HandlerFieldFightLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.FieldFightLoadAck{}
	err = m.FieldFight.LoadInfo(user, ack)
	if err != nil {
		return nil, nil, err
	}

	//logger.Debug("HandlerFieldFightLoadReq ack is %v", ack)
	return ack, nil, nil
}

//刷新对手
func HandlerRefFieldFightRivalUserReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	var err error
	ack := &pb.RefFieldFightRivalUserAck{}
	err = m.FieldFight.RivalUser(user, ack, false)
	if err != nil {
		return nil, nil, err
	}

	//logger.Debug("HandlerRefFieldFightRivalUserReq ack is %v", ack)
	return ack, nil, nil
}

//竞技场挑战次数购买
func HandlerBuyFieldFightChallengeTimesReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFieldFightBuyChallenge)

	var err error
	ack := &pb.BuyFieldFightChallengeTimesAck{}
	err = m.FieldFight.BuyFieldFightChallengeNum(user, op, ack)
	if err != nil {
		return nil, nil, err
	}
	//logger.Debug("HandlerBuyFieldFightChallengeTimesReq ack is %v", ack)

	return ack, op, nil
}

//进入战斗
func HandlerEnterFieldFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterFieldFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error

	err = m.FieldFight.EnterFieldFight(user, int(req.ChallengeUid), int(req.IsBeatBack))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
