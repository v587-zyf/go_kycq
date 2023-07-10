package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

const DarkPalaceBossHpSyncInterval = 200

/*暗殿boss*/
type DarkPalaceBossFight struct {
	*DefaultFight
	boss           *actorPkg.MonsterActor
	hpSyncTime     time.Time //上次gs同步boss血量时间
	lastSyncBossHp int       //上次gs同步boss血量值
	nextReliveTime int64     //复活时间
}

func NewDarkPalaceBossFight(stageId int) (*DarkPalaceBossFight, error) {
	var err error
	fight := &DarkPalaceBossFight{}
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

func (this *DarkPalaceBossFight) OnDie(actor base.Actor, killer base.Actor) {
	if actor.GetType() == base.ActorTypeMonster {
		this.nextReliveTime = time.Now().Unix() + int64(this.boss.MonsterT.ReliveDelay/1000)
		this.OnEnd()
	} else {

		mainActor := this.GetUserMainActor(actor.GetUserId())
		if this.boss.Owner() == actor.GetUserId() {
			//仇恨目标清理
			allDie := this.CheckUserAllDie(mainActor)
			if allDie {
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

func (this *DarkPalaceBossFight) PostDamage(attacker, defender base.Actor, damage int) {
	this.DefaultFight.PostDamage(attacker, defender, damage)
	if this.lastSyncBossHp == this.boss.GetProp().HpNow() {
		return
	}
	timeInterval := time.Now().Sub(this.hpSyncTime).Milliseconds()
	if timeInterval > DarkPalaceBossHpSyncInterval {
		this.lastSyncBossHp = this.boss.GetProp().HpNow()
		this.hpSyncTime = time.Now()
		this.sendFieldBossInfo()
	}
}

func (this *DarkPalaceBossFight) GetBossHpPoint() float64 {
	return float64(this.boss.GetProp().HpNow()) / float64(this.boss.GetProp().Get(pb.PROPERTY_HP))
}

func (this *DarkPalaceBossFight) GetReliveTime() int64 {
	if time.Now().Unix() > this.nextReliveTime {
		return 0
	}
	return this.nextReliveTime
}

func (this *DarkPalaceBossFight) GetBossOwner() (int, int) {
	return this.boss.GetObjId(), this.boss.Owner()
}

func (this *DarkPalaceBossFight) OnPickAll(lastPickObjId int) {
	actors := this.GetUserActors()
	if len(actors) > 0 {
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
		}
		winner := this.GetUserMainActor(this.boss.Owner())
		msg.Winners = []int32{int32(winner.GetUserId())}

		userPickMsg := &pbserver.DarkPalaceBossResult{
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

func (this *DarkPalaceBossFight) OnEnd() {

	this.sendFieldBossInfo()
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_READY, int64(this.boss.MonsterT.ReliveDelay/1000))
	winUserId := this.boss.Owner()
	userPickMsg := &pbserver.DarkPalaceBossResult{
		SendWinner: false,
		Helper:     make(map[int32]int32),
	}
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

func (this *DarkPalaceBossFight) OnRelive(actor base.Actor, reliveType int) {

	if actor.GetType() == base.ActorTypeMonster {

		this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
		this.sendFieldBossInfo()
	}
}

func (this *DarkPalaceBossFight) sendFieldBossInfo() {

	ntf := &pbserver.FsFieldBossInfoNtf{
		StageId:    int32(this.StageConf.Id),
		Hp:         float32(this.GetBossHpPoint()),
		ReliveTime: this.GetReliveTime(),
	}
	net.GetGsConn().SendMessage(ntf)
}
