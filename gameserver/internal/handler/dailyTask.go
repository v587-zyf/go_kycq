package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdDailyTaskLoadReqId, HandlerDailyTaskLoadReq)
	pb.Register(pb.CmdBuyChallengeTimeReqId, HandlerBuyChallengeTimeReq)
	pb.Register(pb.CmdGetAwardReqId, HandlerGetAwardReq)
	pb.Register(pb.CmdResourcesBackGetRewardReqId, HandlerResourcesBackGetRewardReq)
	pb.Register(pb.CmdResourcesBackGetAllRewardReqId, HandlerResourcesBackGetAllRewardReq)
	pb.Register(pb.CmdGetExpReqId, HandlerGetExpReq)
}

func HandlerDailyTaskLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.DailyTaskLoadAck{}

	var err error
	err = m.DailyTask.DailyTaskLoad(user, ack)
	if err != nil {
		return nil, nil, err
	}

	return ack, nil, nil
}

func HandlerBuyChallengeTimeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.BuyChallengeTimeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.BuyChallengeTimeAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDailyTaskBuyChallenge)
	var err error
	err = m.DailyTask.BuyChallengeTime(user, int(req.ActivityId), ack, op)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerGetAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.GetAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.GetAwardAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDailyTaskGetAward)
	var err error
	err = m.DailyTask.GetReward(user, int(req.Id), int(req.Type), ack, op)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerResourcesBackGetRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ResourcesBackGetRewardReq)
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.ResourcesBackGetRewardAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDailyTaskResourcesBackGetReward)
	var err error
	err = m.DailyTask.GetResourcesBackReward(user, int(req.UseIngot), int(req.ActivityId), int(req.BackTimes), ack, op)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerResourcesBackGetAllRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.ResourcesBackGetAllRewardAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDailyTaskResourcesBackAllGetReward)
	var err error
	err = m.DailyTask.GetResourcesBackAllReward(user, ack, op)
	if err != nil {
		return nil, nil, err
	}

	return ack, op, nil
}

func HandlerGetExpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	//req := p.(*pb.GetExpReq)
	//user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.GetExpAck{}
	//var err error
	//err = m.DailyTask.GetExp(user, int(req.ActivityId), ack)
	//if err != nil {
	//	return nil, nil, err
	//}

	return ack, nil, nil
}
