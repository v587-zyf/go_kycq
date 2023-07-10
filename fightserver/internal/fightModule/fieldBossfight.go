package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

const FieldBossHpSyncInterval = 200

/*野外boss*/
type FieldBossFight struct {
	*DefaultFight
	boss           *actorPkg.MonsterActor
	hpSyncTime     time.Time //上次gs同步boss血量时间
	lastSyncBossHp int       //上次gs同步boss血量值
	nextReliveTime int64     //复活时间
}

func NewFieldBossFight(stageId int) (*FieldBossFight, error) {
	var err error
	fieldBossFight := &FieldBossFight{}
	fieldBossFight.DefaultFight, err = NewDefaultFight(stageId, fieldBossFight)
	if err != nil {
		return nil, err
	}
	fieldBossFight.SetLifeTime(-1)
	fieldBossFight.InitMonster()
	//只有一个怪物boss
	for _, v := range fieldBossFight.monsterActors {
		fieldBossFight.boss = v.(*actorPkg.MonsterActor)
	}
	fieldBossFight.Start()
	return fieldBossFight, nil
}

func (this *FieldBossFight) OnDie(actor base.Actor, killer base.Actor) {
	if actor.GetType() == base.ActorTypeMonster {
		this.nextReliveTime = time.Now().Unix() + int64(this.boss.MonsterT.ReliveDelay/1000)
		this.OnEnd()
	} else if actor.GetType() == base.ActorTypeUser {

		mainActor := this.GetUserMainActor(actor.GetUserId())
		if this.boss.Owner() == mainActor.GetUserId() {
			//仇恨目标清理
			allDie := this.CheckUserAllDie(mainActor)
			if allDie {
				this.boss.AddOwner(killer, true)
			}
		}
		this.sendUserDieTime(actor.GetUserId())
	}
}

func (this *FieldBossFight) PostDamage(attacker, defender base.Actor, damage int) {
	this.DefaultFight.PostDamage(attacker, defender, damage)
	if this.lastSyncBossHp == this.boss.GetProp().HpNow() {
		return
	}
	timeInterval := time.Now().Sub(this.hpSyncTime).Milliseconds()
	if timeInterval > FieldBossHpSyncInterval {
		this.lastSyncBossHp = this.boss.GetProp().HpNow()
		this.hpSyncTime = time.Now()
		this.sendFieldBossInfo()
	}
}

func (this *FieldBossFight) GetBossHpPoint() float64 {
	return float64(this.boss.GetProp().HpNow()) / float64(this.boss.GetProp().Get(pb.PROPERTY_HP))
}

func (this *FieldBossFight) GetReliveTime() int64 {
	if time.Now().Unix() > this.nextReliveTime {
		return 0
	}
	return this.nextReliveTime
}

func (this *FieldBossFight) GetBossOwner() (int, int) {
	return this.boss.GetObjId(), this.boss.Owner()
}

func (this *FieldBossFight) OnPickAll(lastPickObjId int) {

	if this.boss.GetProp().HpNow() > 0 {
		return
	}

	//if this.boss.Owner() != lastPickObjId {
	//	logger.Error("玩家拾取道具，但拾取道具的玩家不是归属者,玩家归属：%v,拾取者：%v", this.boss.Owner(), lastPickObjId)
	//	return
	//}

	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
	}
	winner := this.GetUserMainActor(this.boss.Owner())
	msg.Winners = []int32{int32(this.boss.Owner())}

	if winner != nil {
		userPickMsg := &pbserver.FieldBossResult{
			UserPickItems: make(map[int32]*pbserver.ItemUnits),
			SendWinner:    true,
		}
		userPickMsg.UserPickItems[int32(winner.GetUserId())] = &pbserver.ItemUnits{Items: make([]*pbserver.ItemUnit, 0)}
		for k, v := range this.playerPickUp[winner.GetUserId()] {
			userPickMsg.UserPickItems[int32(winner.GetUserId())].Items = append(userPickMsg.UserPickItems[int32(winner.GetUserId())].Items, &pbserver.ItemUnit{
				int32(k), int32(v),
			})
		}
		rb, _ := userPickMsg.Marshal()
		msg.CpData = rb

		net.GetGsConn().SendMessage(msg)
	}

	this.playerPickUp = make(map[int]map[int]int)
}

func (this *FieldBossFight) OnEnd() {

	this.sendFieldBossInfo()
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_READY, int64(this.boss.MonsterT.ReliveDelay/1000))
	actors := this.GetAllUserIds()
	if len(actors) > 0 {
		loser := make([]int32, 0)

		winner := this.GetUserMainActor(this.boss.Owner())
		for _, v := range actors {
			if winner != nil && v != winner.GetUserId() {
				loser = append(loser, int32(v))
			}
		}
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
		}
		msg.Winners = []int32{int32(winner.GetUserId())}
		msg.Losers = loser
		userPickMsg := &pbserver.FieldBossResult{
			SendWinner: false,
		}
		rb, _ := userPickMsg.Marshal()
		msg.CpData = rb
		net.GetGsConn().SendMessage(msg)
	}
}

func (this *FieldBossFight) OnRelive(actor base.Actor, reliveType int) {

	if actor.GetType() == base.ActorTypeMonster {

		this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
		this.sendFieldBossInfo()
	}
}

func (this *FieldBossFight) sendFieldBossInfo() {

	ntf := &pbserver.FsFieldBossInfoNtf{
		StageId:    int32(this.StageConf.Id),
		Hp:         float32(this.GetBossHpPoint()),
		ReliveTime: this.GetReliveTime(),
	}
	net.GetGsConn().SendMessage(ntf)
}

func (this *FieldBossFight) sendUserDieTime(dieUserId int) {

	ntf := &pbserver.FsFieldBossDieUserInfoNtf{
		DieUserId: int32(dieUserId),
		DieTime:   time.Now().Unix(),
	}
	net.GetGsConn().SendMessage(ntf)
}
