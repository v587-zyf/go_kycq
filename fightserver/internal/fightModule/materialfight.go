package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"time"
)

type MaterialFight struct {
	*DefaultFight
	*base.FightReady
	nowWave     int //当前第几波
	fightResult bool
}

func NewMaterialFight(stageId int) (*MaterialFight, error) {
	var err error
	materialFight := &MaterialFight{
		nowWave: 0,
	}
	materialFight.DefaultFight, err = NewDefaultFight(stageId, materialFight)
	if err != nil {
		return nil, err
	}
	//materialFight.InitMonsterByWave(materialFight.nowWave)
	materialFight.FightReady = base.NewFightReady()
	materialFight.Start()

	return materialFight, nil
}

func (this *MaterialFight) OnDie(actor base.Actor, killer base.Actor) {

	if actor.GetType() == base.ActorTypeMonster {

		allDie := true
		for _, monster := range this.monsterActors {
			if monster.GetProp().HpNow() > 0 {
				allDie = false
				break
			}
		}
		if allDie {

			if this.nowWave == len(this.StageConf.Monster_group)-1 {

				this.fightResult = true
				this.OnEnd()
			} else {
				this.nowWave += 1
				//for _, monster := range this.monsterActors {
				//	if monster.GetObjId() != actor.GetObjId() && monster.GetProp().HpNow() <= 0 {
				//		this.Leave(monster)
				//	}
				//}
				this.InitMonsterByWave(this.nowWave)
			}
		}
	} else {
		allDie := this.CheckUserAllDie(actor)
		if allDie {
			this.fightResult = false
			this.OnEnd()
		}
	}
}

func (this *MaterialFight) OnLeaveUser(userId int) {
	this.Stop()
}

func (this *MaterialFight) OnEnd() {
	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	var fightOwner base.Actor
	actors := this.GetUserActors()
	if len(actors) > 0 {
		//材料副本里面肯定只有一个人
		for _, v := range actors {
			if v.HostId() > 0 {
				fightOwner = v
			}
		}
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
		}
		if this.fightResult {
			msg.Winners = []int32{int32(fightOwner.GetUserId())}
		} else {
			msg.Losers = []int32{int32(fightOwner.GetUserId())}
		}
		logger.Info("发送game材料副本战斗结果,服务器：%v，结果：%v", fightOwner.HostId(), *msg)
		net.GetGsConn().SendMessageToGs(uint32(fightOwner.HostId()), msg)
	}
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}

func (this *MaterialFight) Begin() {

	this.InitMonsterByWave(this.nowWave)
	//重置战斗开始时间
	this.createTime = time.Now().Unix()
}
