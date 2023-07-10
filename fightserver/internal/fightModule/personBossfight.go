package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type PersonBossFight struct {
	*DefaultFight
	fightResult int
}

func NewPersonBossFight(stageId int) (*PersonBossFight, error) {
	var err error
	personBossFight := &PersonBossFight{}
	personBossFight.DefaultFight, err = NewDefaultFight(stageId, personBossFight)
	if err != nil {
		return nil, err
	}
	personBossFight.InitMonster()
	personBossFight.Start()

	return personBossFight, nil
}

func (this *PersonBossFight) OnDie(actor base.Actor, killer base.Actor) {
	if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {

		//monster := actor.(*actorPkg.MonsterActor)
		//if monster.MonsterT.Type
		this.fightResult = pb.RESULTFLAG_SUCCESS
		//this.OnEnd()

	} else {
		allDie := this.CheckUserAllDie(actor)
		if allDie {
			this.fightResult = pb.RESULTFLAG_FAIL
			this.OnEnd()
		}
	}
}

func (this *PersonBossFight) OnLeaveUser(userId int) {
	this.Stop()
}

func (this *PersonBossFight) OnPickAll(lastPickObjId int) {
	this.OnEnd()
}

func (this *PersonBossFight) OnEnd() {
	if this.status != FIGHT_STATUS_RUNNING{
		return
	}
	var fightOwner base.Actor
	actors := this.GetUserActors()
	if len(actors) > 0 {
		//个人boss里面肯定只有一个人
		for _, v := range actors {
			if v.HostId() > 0 {
				fightOwner = v
				break
			}
		}
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
		}
		userId := fightOwner.GetUserId()
		resultMsg := &pbserver.PersonFightResult{
			UserId: int32(fightOwner.GetUserId()),
			Result: int32(this.fightResult),
			Items:  make(map[int32]int32),
		}
		for k, v := range this.playerPickUp[userId] {
			resultMsg.Items[int32(k)] = int32(v)
		}
		rb, _ := resultMsg.Marshal()
		msg.CpData = rb
		logger.Info("发送game个人boss战斗结果,服务器：%v，结果：%v", fightOwner.HostId(), *msg)
		net.GetGsConn().SendMessageToGs(uint32(fightOwner.HostId()), msg)
	}
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}
