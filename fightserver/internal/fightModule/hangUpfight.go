package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type HangUpFight struct {
	*DefaultFight
	nowWave     int //当前第几波
	fightResult bool
	fightOwner  base.Actor
}

func NewHangUpFight(stageId int) (*HangUpFight, error) {
	var err error
	f := &HangUpFight{
		nowWave: 0,
	}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}
	f.SetLifeTime(-1)
	f.InitMonsterByWave(f.nowWave)
	f.Start()

	return f, nil
}

func (this *HangUpFight) OnDie(actor base.Actor, killer base.Actor) {

	if actor.GetType() == base.ActorTypeMonster {

		allDie := true
		for _, monster := range this.monsterActors {
			if monster.GetProp().HpNow() > 0 {
				allDie = false
			}
		}
		logger.Debug("当前第几波：%v,全部死亡：%v,玩家：%v,限制：%v", this.nowWave, allDie, actor.GetObjId(), gamedb.GetConf().ExpStageRefreshMonster)

		if allDie {
			if this.nowWave == len(this.StageConf.Monster_group)-1 {

				this.nowWave = 0
			} else {
				this.nowWave += 1
			}
			for _, monster := range this.monsterActors {
				if actor.GetObjId() != monster.GetObjId() {
					this.Leave(monster)
				}
			}

			this.InitMonsterByWave(this.nowWave)
			this.sendUserKillWave()
		}

	} else {
		allDie := this.CheckUserAllDie(actor)
		if allDie {
			this.fightResult = false
			this.OnEnd()
		}
	}
}

func (this *HangUpFight) OnActorEnter(actor base.Actor) {

	this.DefaultFight.OnActorEnter(actor)
	if actor.GetType() == pb.SCENEOBJTYPE_USER {
		if actor.HostId() > 0 {

			this.fightOwner = actor
		}
		actor.AddBuff(gamedb.GetConf().GuajiBuff, actor, false)
	}

}

func (this *HangUpFight) OnLeaveUser(userId int) {
	this.Stop()
}

func (this *HangUpFight) sendUserKillWave() {

	msg := &pbserver.HangUpKillWaveNtf{
		UserId:  int32(this.fightOwner.GetUserId()),
		StageId: int32(this.GetStageConf().Id),
	}
	net.GetGsConn().SendMessageToGs(uint32(this.fightOwner.HostId()), msg)
}
