package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

type ArenaFight struct {
	*DefaultFight
	*base.FightReady
	*base.FightTotal
	fightOwner  base.Actor
	fightResult bool
	cpData      []byte
}

func NewArenaFight(stageId int, cpData []byte) (*ArenaFight, error) {
	var err error
	arenaFight := &ArenaFight{}
	arenaFight.DefaultFight, err = NewDefaultFight(stageId, arenaFight)
	arenaFight.FightReady = base.NewFightReady()
	if err != nil {
		return nil, err
	}
	arenaFight.InitMonster()
	arenaFight.Start()
	arenaFight.cpData = cpData

	return arenaFight, nil
}

func (this *ArenaFight) OnDie(actor base.Actor, killer base.Actor) {
	allDie := this.CheckUserAllDie(actor)
	if allDie {
		if actor.GetUserId() == this.fightOwner.GetUserId() {

			this.fightResult = false
		} else {
			this.fightResult = true
		}
		this.OnEnd()
	}
}

func (this *ArenaFight) OnActorEnter(actor base.Actor) {

	if actor.SessionId() != 0 {
		this.fightOwner = actor
		this.FightReady.FightInto(actor.GetUserId())
	}
	if actor.GetType() == pb.SCENEOBJTYPE_FIT {
		actor.GetFSM().Recover()
	}
	if actor.GetType() == pb.SCENEOBJTYPE_SUMMON && actor.GetUserId() != this.fightOwner.GetUserId() {
		actor.GetFSM().Recover()
	}
}

func (this *ArenaFight) Begin() {

	for _, v := range this.userActors {
		v.GetFSM().Recover()
	}
	for _, v := range this.userPets {
		if v.TeamIndex() == constFight.FIGHT_TEAM_ZERO {
			v.GetFSM().Recover()
		}
	}

	for _, v := range this.userSummons {
		if v.TeamIndex() == constFight.FIGHT_TEAM_ZERO {
			v.GetFSM().Recover()
		}
	}
	//重置战斗开始时间
	this.createTime = time.Now().Unix()
}

func (this *ArenaFight) OnLeaveUser(userId int) {
	//if actor.GetType() == base.ActorTypeUser {
	this.Stop()
	//	msg := &pbserver.FSFightEndNtf{
	//		FightType: int32(this.StageConf.Type),
	//		StageId:   int32(this.StageConf.Id),
	//		CpData:    this.cpData,
	//	}
	logger.Info("发送game竞技场战斗结果,玩家离开 服务器：%v，结果：%v", userId)
	//	net.GetGsConn().SendMessageToGs(uint32(actor.HostId()), msg)
	//}
}

func (this *ArenaFight) OnEnd() {
	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		CpData:    this.cpData,
	}
	if this.fightResult {
		msg.Winners = []int32{int32(this.fightOwner.GetUserId())}
	} else {
		msg.Losers = []int32{int32(this.fightOwner.GetUserId())}
	}
	logger.Info("发送game竞技场战斗结果,服务器：%v，结果：%v", this.fightOwner.HostId(), *msg)
	net.GetGsConn().SendMessageToGs(uint32(this.fightOwner.HostId()), msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}
