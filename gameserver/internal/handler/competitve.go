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
	pb.Register(pb.CmdCompetitveLoadReqId, HandlerCompetitveLoadReq)
	pb.Register(pb.CmdBuyCompetitveChallengeTimesReqId, HandlerBuyCompetitveChallengeTimesReq)
	pb.Register(pb.CmdRefCompetitveRankReqId, HandlerRefCompetitveRankReq)
	pb.Register(pb.CmdGetCompetitveDailyRewardReqId, HandlerGetCompetitveDailyRewardReq)
	pb.Register(pb.CmdEnterCompetitveFightReqId, HandlerEnterCompetitveFightReq)
	pb.Register(pb.CmdCompetitveMultipleClaimReqId, HandlerCompetitveMultipleClaimReq)

}

//打开页面
func HandlerCompetitveLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.CompetitveLoadAck{}
	err := m.Competitve.LoadInfo(user, ack)
	if err != nil {
		return nil, nil, err
	}

	logger.Debug("HandlerCompetitveLoadReq ack is %v", ack)
	return ack, nil, nil
}

//竞技场挑战次数购买
func HandlerBuyCompetitveChallengeTimesReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCompetitveBuyChallenge)

	var err error
	ack := &pb.BuyCompetitveChallengeTimesAck{}
	err = m.Competitve.BuyCompetitiveChallengeNum(user, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerBuyArenaFightNumReq ack is %v", ack)

	return ack, op, nil
}

func HandlerGetCompetitveDailyRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCompetitveDailyReward)

	var err error
	ack := &pb.GetCompetitveDailyRewardAck{}

	err = m.Competitve.GetCompetitiveDailyReward(user, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerBuyArenaFightNumReq ack is %v", ack)

	return ack, op, nil
}

//刷新对手
func HandlerRefCompetitveRankReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	var err error
	ack := &pb.RefCompetitveRankAck{}

	err = m.Competitve.RefCompetitiveRivalNew(user, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerRefArenaRankReq ack is %v", ack)

	return ack, nil, nil
}

//进入战斗
func HandlerEnterCompetitveFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterCompetitveFightReq)
	user := conn.GetSession().(*managers.ClientSession).User

	var err error

	err = m.Competitve.EnterCompetitveFight(user, int(req.ChallengeUid), int(req.ChallengeRanking))
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

//竞技场多倍领取
func HandlerCompetitveMultipleClaimReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCompetitveMultipleClaim)

	var err error
	ack := &pb.CompetitveMultipleClaimAck{}
	err = m.Competitve.CompetitiveMultipleClaim(user, op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerBuyArenaFightNumReq ack is %v", ack)

	return ack, op, nil
}
