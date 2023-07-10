package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

const HellBossHpSyncInterval = 200

/*炼狱boss*/
type HellBossFight struct {
	*DefaultFight
	boss           *actorPkg.MonsterActor
	hpSyncTime     time.Time //上次gs同步boss血量时间
	lastSyncBossHp int       //上次gs同步boss血量值
	nextReliveTime int64     //复活时间
}

func NewHellBossFight(stageId int) (*HellBossFight, error) {
	var err error
	fight := &HellBossFight{}
	fight.DefaultFight, err = NewDefaultFight(stageId, fight)
	if err != nil {
		return nil, err
	}
	fight.SetLifeTime(-1)
	fight.InitMonster()
	//只有一个怪物boss
	for _, v := range fight.monsterActors {
		fight.boss = v.(*actorPkg.MonsterActor)
	}
	fight.Start()
	return fight, nil
}

func (this *HellBossFight) OnDie(actor base.Actor, killer base.Actor) {
	if actor.GetType() == base.ActorTypeMonster {
		this.nextReliveTime = time.Now().Unix() + int64(this.boss.MonsterT.ReliveDelay/1000)
		this.OnEnd()
	} else {
		userId := actor.GetUserId()
		mainActor := this.GetUserMainActor(userId)
		if this.boss.Owner() == userId {
			if allDie := this.CheckUserAllDie(mainActor); allDie {
				killerUserId := killer.GetUserId()
				if killerUserId > 0 {
					killerMainActor := this.GetUserMainActor(killerUserId)
					if killerMainActor != nil && killerMainActor.(base.ActorPlayer).GetPlayer().FightNum() > 0 {
						this.boss.AddOwner(killer, true)
					} else {
						this.boss.AddOwner(nil, true)
					}
				} else {
					this.boss.AddOwner(nil, true)
				}
			}
		}
	}
}

func (this *HellBossFight) PostDamage(attacker, defender base.Actor, damage int) {
	this.DefaultFight.PostDamage(attacker, defender, damage)
	if this.lastSyncBossHp == this.boss.GetProp().HpNow() {
		return
	}
	if timeInterval := time.Now().Sub(this.hpSyncTime).Milliseconds(); timeInterval > HellBossHpSyncInterval {
		this.lastSyncBossHp = this.boss.GetProp().HpNow()
		this.hpSyncTime = time.Now()
		this.sendBossInfo()
	}
}

func (this *HellBossFight) OnPickAll(lastPickObjId int) {
	actors := this.GetUserActors()
	if len(actors) > 0 {
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
		}
		winner := this.GetUserMainActor(this.boss.Owner())
		winnerUserId := int32(winner.GetUserId())
		msg.Winners = []int32{winnerUserId}

		userPickMsg := &pbserver.HellBossResult{
			UserPickItems: make(map[int32]*pbserver.ItemUnits),
			SendWinner:    true,
		}
		userPickMsg.UserPickItems[winnerUserId] = &pbserver.ItemUnits{}
		pickItemSlice := make([]*pbserver.ItemUnit, 0)
		for k, v := range this.playerPickUp[int(winnerUserId)] {
			pickItemSlice = append(pickItemSlice, &pbserver.ItemUnit{ItemId: int32(k), ItemNum: int32(v)})
		}
		userPickMsg.UserPickItems[winnerUserId].Items = pickItemSlice
		rb, _ := userPickMsg.Marshal()
		msg.CpData = rb

		net.GetGsConn().SendMessage(msg)
	}
	this.playerPickUp = make(map[int]map[int]int)
}

func (this *HellBossFight) OnLeave(actor base.Actor) {
	if actor.GetType() == base.ActorTypeMonster {
		this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
		this.sendBossInfo()
	}
}

func (this *HellBossFight) OnEnd() {
	this.sendBossInfo()
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_READY, int64(this.boss.MonsterT.ReliveDelay/1000))
	winUserId := this.boss.Owner()
	userPickMsg := &pbserver.HellBossResult{Helper: make(map[int32]int32)}
	loser := make([]int32, 0)
	for k, v := range this.playerActors {
		if k != winUserId {
			if v.ToHelpUserId() > 0 {
				userPickMsg.Helper[int32(k)] = int32(v.ToHelpUserId())
			} else {
				loser = append(loser, int32(k))
			}
		}
	}
	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
	}
	msg.Winners = []int32{int32(winUserId)}
	msg.Losers = loser

	rb, _ := userPickMsg.Marshal()
	msg.CpData = rb
	net.GetGsConn().SendMessage(msg)
}

func (this *HellBossFight) OnRelive(actor base.Actor, reliveType int) {

	if actor.GetType() == base.ActorTypeMonster {
		this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
		this.sendBossInfo()
	}
}

func (this *HellBossFight) GetBossHpPoint() float64 {
	return float64(this.boss.GetProp().HpNow()) / float64(this.boss.GetProp().Get(pb.PROPERTY_HP))
}

func (this *HellBossFight) GetReliveTime() int64 {
	if time.Now().Unix() > this.nextReliveTime {
		return 0
	}
	return this.nextReliveTime
}

func (this *HellBossFight) GetBossOwner() (int, int) {
	return this.boss.GetObjId(), this.boss.Owner()
}

func (this *HellBossFight) sendBossInfo() {
	ntf := &pbserver.FsFieldBossInfoNtf{
		StageId:    int32(this.StageConf.Id),
		Hp:         float32(this.GetBossHpPoint()),
		ReliveTime: this.GetReliveTime(),
	}
	net.GetGsConn().SendMessage(ntf)
}
