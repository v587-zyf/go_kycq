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
	pb.Register(pb.CmdStageFightStartReqId, HandlerStageFightStartReq)
	//pb.Register(pb.CmdStageFightEndReqId, HandlerStageFightEndReq)
	pb.Register(pb.CmdLeaveFightReqId, HandlerLeaveFightReq)
	//pb.Register(pb.CmdKillMonsterReqId, HandlerKillMonsterReq)
	pb.Register(pb.CmdStartStageBossFightReqId, HandlerStartStageBossFight)
	pb.Register(pb.CmdStageGetHookMapRewardReqId, HandlerGetHookMapRewards)
}

func HandlerStageFightStartReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.StageFightStartReq)
	//if int(req.StageId) != user.StageId {
	//	user.StageId = int(req.StageId)
	//	user.StageWave = 0
	//}
	err := m.StageManager.StageFightStartReq(user)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("接收到客户端发来的挂机开始：客户端数据：%v--%v，玩家当前为：%v--%v", req.StageId, req.Wave, user.StageId, user.StageWave)
	return &pb.StageFightStartAck{
		StageId: int32(user.StageId),
		Wave:    int32(user.StageWave),
	}, nil, nil
}

func HandlerStageFightEndReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeNormalStage)
	err := m.StageManager.StageFightEndReq(user, op)
	if err != nil {
		logger.Error("关卡通关异常：%v", err)
		return nil, nil, err
	}
	return &pb.StageFightEndNtf{StageId: int32(user.StageId), Wave: int32(user.StageWave), Goods: op.ToChangeItems()}, op, nil
}

func HandlerLeaveFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Fight.LeaveFight(user, constFight.LEAVE_FIGHT_TYPE_NOMAL)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.LeaveFightAck{}
	return ack, nil, nil
}

func HandlerKillMonsterReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.KillMonsterReq)
	user := conn.GetSession().(*managers.ClientSession).User
	m.GetTask().UpdateTaskForKillMonster(user, int(req.MonsterId), int(req.KillNum))
	ack := &pb.KillMonsterAck{}
	return ack, nil, nil
}

func HandlerStartStageBossFight(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.StageManager.StartStageBossFight(user)
	return nil, nil, err
}

func HandlerGetHookMapRewards(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeNormalStage)
	err := m.StageManager.GetStageHookMapReward(user, op)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.StageGetHookMapRewardAck{
		Goods:      op.ToChangeItems(),
		Items:      make([]*pb.ItemUnit, 0),
		HookupTime: 0,
	}
	m.GetUserManager().SendItemChangeNtf(user, op)
	return ack, nil, nil
}
