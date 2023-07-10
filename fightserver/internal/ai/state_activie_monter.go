package ai

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/fsm"
	"cqserver/golibs/common"
	"math/rand"
	"time"
)

type MonsterActiveState struct {
	*ActivieState
	randMoveTime int64
}

func NewMonsterActiveState(aifsm *fsm.FSM, actor base.Actor) *MonsterActiveState {
	return &MonsterActiveState{
		ActivieState: &ActivieState{
			actor:    actor,
			fsm:      aifsm,
			findTime: time.Now().Unix(),
		},
	}
}

//func (this *MonsterActiveState) ExecuteBefore() {
//	this.ActivieState.ExecuteBefore()
//}

func (this *MonsterActiveState) Enter() {

	this.randMoveTime = 0
	//清空路径
	this.actor.SetPathSlice(nil)
}

func (this *MonsterActiveState) Leave() {
	//清空路径
	this.actor.SetPathSlice(nil)
}

func (this *MonsterActiveState) Execute() {

	this.ActivieState.Execute()
	//判断怪物活动区域，在活动区域内，查找获取区域最近的目标，不在活动区域内，则寻路回家
	switch this.actor.(*actorPkg.MonsterActor).MonsterT.Aitype {

	case 1:
		isFinder := this.FindEnemyFor(false)
		if isFinder {
			return
		}
	default:
		return
	}
	isMove := true
	if this.actor.(*actorPkg.MonsterActor).HasTheadTarget() {
		this.fsm.Event(base.StateTriggerFight)
		isMove = false
	}

	if isMove {
		this.move()
	}
}

//主动怪物
func (this *MonsterActiveState) FindEnemyFor(isRand bool) bool {

	monsterActor := this.actor.(*actorPkg.MonsterActor)
	birthPoint := monsterActor.BirthPoint()
	curPoint := monsterActor.Point()

	disToBirth := scene.Distance(birthPoint.X(), birthPoint.Y(), curPoint.X(), curPoint.Y())

	var enemy base.Actor
	if common.GetNowMillisecond()-this.findTime > 0 {
		if disToBirth <= this.actor.(*actorPkg.MonsterActor).MonsterT.Toattackarea {
			//巡逻范围，查找最近的敌人
			enemy = findNearestMonsterEnemy(this.actor, this.actor.(*actorPkg.MonsterActor).MonsterT.Toattackarea)
		}
		this.findTime = common.GetNowMillisecond() + 2000
	}

	if enemy != nil {
		monsterActor.AddTheat(enemy.GetObjId(), 0)
		this.fsm.Event(base.StateTriggerFight)
		return true
	}
	return false
}

func (this *MonsterActiveState) move() {
	monsterActor := this.actor.(*actorPkg.MonsterActor)
	if monsterActor.MonsterT.Move == 1 {
		if !this.actor.CanMove() {
			return
		}

		birthPoint := monsterActor.BirthPoint()
		curPoint := monsterActor.Point()
		disToBirth := scene.Distance(birthPoint.X(), birthPoint.Y(), curPoint.X(), curPoint.Y())

		if disToBirth > monsterActor.MonsterT.Toattackarea { // 超过警戒范围的一半，就往出生点走
			monsterHomeWalk(monsterActor)
		} else { // 没超过警戒范围一半，就随机走
			if common.GetNowMillisecond() > this.randMoveTime {
				monsterRandWalk(monsterActor)
				this.randMoveTime = common.GetNowMillisecond() + 3000 + int64(rand.Intn(2000))
			}
		}
	}
	this.fsm.SetSleepTime(1000)
}
