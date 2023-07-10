package handler

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"time"
)

func init() {
	pb.Register(pb.CmdOpenTowerReqId, HandlerOpenTowerReq)
	pb.Register(pb.CmdEnterTowerFightReqId, HandlerEnterTowerFightReq)
	pb.Register(pb.CmdTowerFightContinueReqId, HandlerTowerFightContinueReq)
	pb.Register(pb.CmdToweryDayAwardReqId, HandlerToweryDayAwardReq)
	pb.Register(pb.CmdTowerLotteryReqId, HandlerTowerLotteryReq)
	pb.Register(pb.CmdTowerSweepReqId, HandlerTowerSweepReq)
	//pb.Register(pb.CmdTowerRankRewardReqId, HandlerTowerRankRewardReq)
}

func HandlerOpenTowerReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	timeNow := time.Now()
	day := common.GetResetTime(timeNow)
	week := common.GetYearWeek(timeNow)
	if user.Tower == nil {
		user.Tower = &model.Tower{
			TowerLv: 1,
			Lottery: make([]int, 0),
		}
	}
	userTower := user.Tower
	if userTower.TowerLv <= 0 {
		userTower.TowerLv = 1
	}
	user.Dirty = true
	return &pb.OpenTowerAck{
		TowerLv:    int32(userTower.TowerLv),
		DayAward:   userTower.DayAwardState == day,
		LotteryNum: int32(userTower.LotteryNum),
		LotterId:   int32(userTower.LotteryId),
		RankReward: user.Tower.RankAwardTime == week,
	}, nil, nil
}

func HandlerEnterTowerFightReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Tower.EnterTowerFight(user)
	if err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}

func HandlerTowerFightContinueReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	if err := m.Tower.TowerFightContinue(user); err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}

func HandlerToweryDayAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTowerDayAward)
	if err := m.Tower.DayAward(user, op); err != nil {
		return nil, nil, err
	}
	return &pb.ToweryDayAwardAck{DayAward: true, Goods: op.ToChangeItems()}, op, nil
}

func HandlerTowerLotteryReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTowerLottery)
	if _, err := m.Tower.Lottery(user, op); err != nil {
		return nil, nil, err
	}
	return &pb.TowerLotteryAck{LotteryNum: int32(user.Tower.LotteryNum), LotteryId: int32(user.Tower.LotteryId), Goods: op.ToChangeItems()}, op, nil
}

func HandlerTowerSweepReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTowerSweep)
	ack := &pb.TowerSweepAck{}
	if err := m.Tower.TowerSweep(user, op, ack); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

//func HandlerTowerRankRewardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
//	user := conn.GetSession().(*managers.ClientSession).User
//	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTowerRankReward)
//	if err := m.Tower.RankReward(user, op); err != nil {
//		return nil, nil, err
//	}
//	return &pb.TowerRankRewardAck{RankReward: true, Goods: op.ToChangeItems()}, op, nil
//}
