package base

import (
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

type IFightReady interface {
	FightEnterOk(actor Actor)
	FightStartCountDownOk(actor Actor)
	FightInto(userId int)
}

const (
	READY_STATE_ENTER     = 1
	READY_STATE_COUNTDOWN = 2
)

//
//  @Description: 战斗准备类
//
type FightReady struct {
	users map[int]int
}

func NewFightReady() *FightReady {
	fightReady := &FightReady{
		users: make(map[int]int),
	}
	return fightReady
}

func (fd *FightReady) FightInto(userId int) {
	if userId != 0 {
		fd.users[userId] = 0
	}
}

func (fd *FightReady) FightEnterOk(actor Actor) {

	userId := actor.GetUserId()
	fd.users[userId] = READY_STATE_ENTER
	logger.Debug("场景玩家状态：%v",fd.users)
	allEnterOk := true
	for _, v := range fd.users {
		if v != READY_STATE_ENTER {
			allEnterOk = false
		}
	}
	if allEnterOk {
		actor.NotifyNearby(actor, &pb.FightStartCountDownNtf{
			ServerTime:    int32(common.GetTimeSeconds(time.Now())),
			CountDownTime: int32(common.GetTimeSeconds(time.Now()) + 3),
		}, nil)
	}

}
func (fd *FightReady) FightStartCountDownOk(actor Actor) {

	userId := actor.GetUserId()
	fd.users[userId] = READY_STATE_COUNTDOWN
	logger.Debug("场景玩家倒计时状态：%v",fd.users)
	allEnterOk := true
	for _, v := range fd.users {
		if v != READY_STATE_COUNTDOWN {
			allEnterOk = false
		}
	}
	if allEnterOk {
		actor.NotifyNearby(actor, &pb.FightStartNtf{}, nil)
	}
	actor.GetFight().Begin()

}
