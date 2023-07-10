package ai

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/fsm"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type PlayerDieState struct {
	fsm.DefaultState
	actor base.Actor
	fsm   *fsm.FSM
	isDie bool
}

func NewPlayerDeathState(aifsm *fsm.FSM, actor base.Actor) *PlayerDieState {
	return &PlayerDieState{
		actor: actor,
		fsm:   aifsm,
		isDie: false,
	}
}

func (this *PlayerDieState) Enter() {

	killNtfToGs(this.actor.Killer(), this.actor)

	this.actor.OnDie()
	if this.actor.DeathReason() != constFight.DEATH_REASON_FIT {
		this.actor.GetFight().OnDie(this.actor, this.actor.Killer())
	}
}

func (this *PlayerDieState) Execute() {

	if this.actor.GetType() == pb.SCENEOBJTYPE_USER {
		actorUser := this.actor.(*actorPkg.UserActor)
		if actorUser.ReliveSelf() > 0 && common.GetNowMillisecond() > actorUser.ReliveSelf() {
			this.actor.Relive(constFight.RELIVE_ADDR_TYPE_SITU, constFight.RELIVE_TYPE_BUFF)
		}
	}
}

func killNtfToGs(killer base.Actor, beKiller base.Actor) {

	if killer == nil || killer.GetFight() == nil {
		return
	}
	killerUserId := killer.GetUserId()
	if killerUserId <= 0 {
		return
	}
	killerMainUser := killer.GetFight().GetUserMainActor(killerUserId)
	if killerMainUser == nil {
		return
	}

	if killerMainUser.(*actorPkg.UserActor).UserType() != constFight.FIGHT_USER_TYPE_PLAYER {
		return
	}

	isAllDie := true
	isPlayer := false
	beKillerId := 0
	if beKiller.GetType() == pb.SCENEOBJTYPE_USER || beKiller.GetType() == pb.SCENEOBJTYPE_FIT {
		isAllDie = killer.GetFight().CheckUserAllDieByHp(beKiller)
		isPlayer = true
		beKillerId = beKiller.GetUserId()
	} else if beKiller.GetType() == pb.SCENEOBJTYPE_MONSTER {
		beKillerId = beKiller.(*actorPkg.MonsterActor).MonsterT.Monsterid
	} else {
		return
	}

	if !isAllDie {
		return
	}
	Ntf := &pbserver.FsToGsActorKillNtf{
		Killer:   int32(killerUserId),
		BeKiller: int32(beKillerId),
		IsPlayer: isPlayer,
	}
	net.GetGsConn().SendMessage(Ntf)
}
