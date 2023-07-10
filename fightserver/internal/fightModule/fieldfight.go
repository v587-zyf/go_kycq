package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"strconv"
	"time"
)

type FieldFight struct {
	*DefaultFight
	*base.FightReady
	fightOwner  int
	enmyUser    int
	fightResult bool
	cpData      []byte
}

func NewFieldFight(stageId int, cpData []byte) (*FieldFight, error) {
	var err error
	fieldFight := &FieldFight{}
	fieldFight.DefaultFight, err = NewDefaultFight(stageId, fieldFight)
	fieldFight.FightReady = base.NewFightReady()
	if err != nil {
		return nil, err
	}
	fieldFight.InitMonster()
	fieldFight.Start()
	fieldFight.cpData = cpData

	return fieldFight, nil
}

func (this *FieldFight) OnDie(actor base.Actor, killer base.Actor) {

	allDie := this.CheckUserAllDie(actor)
	if allDie {
		if actor.GetUserId() == this.fightOwner {

			this.fightResult = false
		} else {
			this.fightResult = true
		}
		this.OnEnd()
	}
}

func (this *FieldFight) OnActorEnter(actor base.Actor) {

	if actor.SessionId() != 0 {
		this.fightOwner = actor.GetUserId()
		this.FightReady.FightInto(actor.GetUserId())
	} else {
		if actor.TeamIndex() == constFight.FIGHT_TEAM_ZERO {

			this.enmyUser = actor.GetUserId()
			if u, ok := actor.(*actorPkg.UserActor); ok {
				userType := u.UserType()
				if userType == constFight.FIGHT_USER_TYPE_CONF {
					this.enmyUser = -this.enmyUser
				}
			}
		}
	}
	if actor.GetType() == pb.SCENEOBJTYPE_FIT {
		actor.GetFSM().Recover()
	}
	if actor.GetType() == pb.SCENEOBJTYPE_SUMMON && actor.GetUserId() != this.fightOwner {
		actor.GetFSM().Recover()
	}
}

func (this *FieldFight) Begin() {

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
	this.createTime = time.Now().Unix()
}

func (this *FieldFight) OnLeaveUser(userId int) {
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

func (this *FieldFight) OnEnd() {
	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	owner := this.GetUserMainActor(this.fightOwner)
	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		CpData:    this.cpData,
	}
	if this.fightResult {
		msg.Winners = []int32{int32(owner.GetUserId())}
	} else {
		msg.Losers = []int32{int32(owner.GetUserId())}
	}
	isBeatBack, _ := strconv.Atoi(string(this.cpData))
	msg.CpData = []byte(fmt.Sprintf("%d,%d", this.enmyUser, isBeatBack))
	logger.Info("发送game野战战斗结果,服务器：%v，结果：%v", owner.HostId(), *msg)
	net.GetGsConn().SendMessageToGs(uint32(owner.HostId()), msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}
