package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type VipBossFight struct {
	*DefaultFight
	fightResult int
}

func NewVipBossFight(stageId int) (*VipBossFight, error) {
	var err error
	vipBossFight := &VipBossFight{
		fightResult: -1,
	}
	vipBossFight.DefaultFight, err = NewDefaultFight(stageId, vipBossFight)
	if err != nil {
		return nil, err
	}
	vipBossFight.InitMonster()
	vipBossFight.Start()

	return vipBossFight, nil
}

func (this *VipBossFight) OnDie(actor base.Actor, killer base.Actor) {
	if actor.GetType() == base.ActorTypeUser {
		allDie := this.CheckUserAllDie(actor)
		if allDie {
			this.fightResult = pb.RESULTFLAG_FAIL
			this.OnEnd()
		}
	} else {
		allDie := true
		for _, monster := range this.monsterActors {
			if monster.GetProp().HpNow() > 0 {
				allDie = false
				break
			}
		}
		if allDie {
			this.fightResult = pb.RESULTFLAG_SUCCESS
			//this.OnEnd()
		}
	}
}

func (this *VipBossFight) OnLeaveUser(userId int) {
	this.Stop()
}

func (this *VipBossFight) OnPickAll(lastPickObjId int) {
	if this.fightResult >= 0 {
		this.OnEnd()
	}
}

func (this *VipBossFight) OnEnd() {
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
			}
		}
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
		}
		userId := fightOwner.GetUserId()
		resultMsg := &pbserver.VipBossFightResult{
			UserId: int32(fightOwner.GetUserId()),
			Result: int32(this.fightResult),
			Items:  make(map[int32]int32),
		}
		for k, v := range this.playerPickUp[userId] {
			resultMsg.Items[int32(k)] = int32(v)
		}
		rb, _ := resultMsg.Marshal()
		msg.CpData = rb
		logger.Info("发送game vip boss战斗结果,服务器：%v，消息：%v，结果：%v", fightOwner.HostId(), *msg, this.fightResult)
		net.GetGsConn().SendMessageToGs(uint32(fightOwner.HostId()), msg)
	}
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}
