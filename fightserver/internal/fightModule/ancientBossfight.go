package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

const AncientBossHpSyncInterval = 200

/*远古首领*/
type AncientBossFight struct {
	*DefaultFight
	boss           *actorPkg.MonsterActor
	hpSyncTime     time.Time //上次gs同步boss血量时间
	lastSyncBossHp int       //上次gs同步boss血量值
	nextReliveTime int64     //复活时间
	owners         []int     //最近的归属者
}

func NewAncientBossFight(stageId int) (*AncientBossFight, error) {
	var err error
	fight := &AncientBossFight{owners: make([]int, 5)}
	fight.DefaultFight, err = NewDefaultFight(stageId, fight)
	if err != nil {
		return nil, err
	}
	fight.SetLifeTime(-1)
	fight.InitMonster()
	//只有一个boss
	for _, v := range fight.monsterActors {
		fight.boss = v.(*actorPkg.MonsterActor)
	}
	fight.Start()
	return fight, nil
}

func (this *AncientBossFight) OnDie(actor base.Actor, killer base.Actor) {
	if actor.GetType() == base.ActorTypeMonster {
		this.nextReliveTime = time.Now().Unix() + int64(this.boss.MonsterT.ReliveDelay/1000)
		this.OnEnd()
	} else if actor.GetType() == base.ActorTypeUser {
		mainActor := this.GetUserMainActor(actor.GetUserId())
		if this.boss.Owner() == mainActor.GetUserId() {
			//仇恨清空
			allDie := this.CheckUserAllDie(mainActor)
			if allDie {
				this.boss.AddOwner(killer, true)
			}
		}
	}
}

func (this *AncientBossFight) OnPickAll(lastPickObjId int) {
	if this.boss.GetProp().HpNow() > 0 {
		return
	}
	//if this.boss.Owner() != lastPickObjId {
	//	return
	//}
	winner := this.GetUserMainActor(this.boss.Owner())
	if winner != nil {
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
			Winners:   []int32{int32(winner.GetUserId())},
		}

		userPickMsg := &pbserver.AncientBossResult{
			SendWinner:    true,
			UserPickItems: make(map[int32]*pbserver.ItemUnits),
		}
		userPickMsg.UserPickItems[int32(winner.GetUserId())] = &pbserver.ItemUnits{Items: make([]*pbserver.ItemUnit, 0)}
		for k, v := range this.playerPickUp[winner.GetUserId()] {
			userPickMsg.UserPickItems[int32(winner.GetUserId())].Items = append(userPickMsg.UserPickItems[int32(winner.GetUserId())].Items, &pbserver.ItemUnit{
				ItemId:  int32(k),
				ItemNum: int32(v),
			})
		}
		rb, _ := userPickMsg.Marshal()
		msg.CpData = rb

		net.GetGsConn().SendMessage(msg)
	}

	this.playerPickUp = make(map[int]map[int]int)
}

func (this *AncientBossFight) OnEnd() {
	this.sendAncientBossInfo()
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
			Winners:   []int32{int32(winner.GetUserId())},
			Losers:    loser,
		}
		userPickMsg := &pbserver.AncientBossResult{
			SendWinner: false,
		}
		rb, _ := userPickMsg.Marshal()
		msg.CpData = rb
		net.GetGsConn().SendMessage(msg)
	}
}

func (this *AncientBossFight) OnRelive(actor base.Actor, reliveType int) {
	if actor.GetType() == base.ActorTypeMonster {
		this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
		this.sendAncientBossInfo()
	}
}

func (this *AncientBossFight) PostDamage(attacker, defender base.Actor, damage int) {
	this.DefaultFight.PostDamage(attacker, defender, damage)
	if this.lastSyncBossHp == this.boss.GetProp().HpNow() {
		return
	}
	timeNow := time.Now()
	timeInterval := timeNow.Sub(this.hpSyncTime).Milliseconds()
	if timeInterval > AncientBossHpSyncInterval {
		this.lastSyncBossHp = this.boss.GetProp().HpNow()
		this.hpSyncTime = timeNow
		this.sendAncientBossInfo()
	}
}

func (this *AncientBossFight) GetBossHpPoint() float64 {
	return float64(this.boss.GetProp().HpNow()) / float64(this.boss.GetProp().Get(pb.PROPERTY_HP))
}

func (this *AncientBossFight) GetReliveTime() int64 {
	if time.Now().Unix() > this.nextReliveTime {
		return 0
	}
	return this.nextReliveTime
}

func (this *AncientBossFight) GetBossOwner() (int, int) {
	return this.boss.GetObjId(), this.boss.Owner()
}

func (this *AncientBossFight) OnEnterUser(userId int) {
	this.sendAncientBossInfo()
}

func (this *AncientBossFight) OnLeaveUser(userId int) {
	this.sendAncientBossInfo()
}

func (this *AncientBossFight) sendAncientBossInfo() {
	ntf := &pbserver.FsFieldBossInfoNtf{
		StageId:    int32(this.StageConf.Id),
		Hp:         float32(this.GetBossHpPoint()),
		ReliveTime: this.GetReliveTime(),
		UserCount:  int32(len(this.GetPlayerUserids())),
	}
	net.GetGsConn().SendMessage(ntf)
}
