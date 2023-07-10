package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

type HangUpBossFight struct {
	*DefaultFight
	*base.FightReady
	fightResult int
	fightOwner  base.Actor
}

func NewHangUpBossFight(stageId int) (*HangUpBossFight, error) {
	var err error
	f := &HangUpBossFight{}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	f.FightReady = base.NewFightReady()
	if err != nil {
		return nil, err
	}
	f.Start()
	f.InitMonster()
	return f, nil
}

func (this *HangUpBossFight) OnDie(actor base.Actor, killer base.Actor) {

	if actor.GetType() == base.ActorTypeMonster {

		allDie := true
		for _, monster := range this.monsterActors {
			if monster.GetProp().HpNow() > 0 {
				allDie = false
			}
		}
		if allDie {
			this.fightResult = pb.RESULTFLAG_SUCCESS
			//this.OnEnd()
		}

	} else {
		allDie := this.CheckUserAllDie(actor)
		if allDie {
			this.fightResult = pb.RESULTFLAG_FAIL
			this.OnEnd()
		}
	}
}

func (this *HangUpBossFight) OnPickAll(lastPickObjId int) {
	this.OnEnd()
}

func (this *HangUpBossFight) OnActorEnter(actor base.Actor) {

	//this.DefaultFight.OnActorEnter(actor)
	if actor.GetType() == pb.SCENEOBJTYPE_USER && actor.HostId() > 0 {
		this.fightOwner = actor
	}
	if actor.GetType() == pb.SCENEOBJTYPE_FIT {
		actor.GetFSM().Recover()
	}
}

func (this *HangUpBossFight) Begin() {

	for _, v := range this.monsterActors {
		v.GetFSM().Recover()
	}
	for _, v := range this.userActors {
		v.GetFSM().Recover()
	}
	this.createTime = time.Now().Unix()
}

func (this *HangUpBossFight) OnLeaveUser(userId int) {
	this.Stop()
}

func (this *HangUpBossFight) OnEnd() {
	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	//挂机boss里面肯定只有一个人
	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
	}
	resultMsg := &pbserver.HangUpBossFightEndNtf{
		UserId:  int32(this.fightOwner.GetUserId()),
		StageId: int32(this.GetStageConf().Id),
		Result:  int32(this.fightResult),
		Items:   make(map[int32]int32),
	}
	for k, v := range this.playerPickUp[this.fightOwner.GetUserId()] {
		resultMsg.Items[int32(k)] = int32(v)
	}
	rb, _ := resultMsg.Marshal()
	msg.CpData = rb
	logger.Debug("发送game个人boss战斗结果,服务器：%v，结果：%v", this.fightOwner.HostId(), *msg)
	net.GetGsConn().SendMessageToGs(uint32(this.fightOwner.HostId()), msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}
