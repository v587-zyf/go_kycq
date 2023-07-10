package ai

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/protobuf/pb"
	"math/rand"
	"time"

	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/fsm"
)

type FightState struct {
	fsm.DefaultState
	actor        base.Actor
	fsm          *fsm.FSM
	lastMoveTime int64
	stepIndex    int
	stepArr      []*scene.Point
	context      IFightStageInterface
}

const (
	FIND_PATH_NUM = 10
)

func NewFightState(aifsm *fsm.FSM, actor base.Actor, ctx IFightStageInterface) *FightState {
	return &FightState{
		actor:   actor,
		fsm:     aifsm,
		context: ctx,
	}
}

func (this *FightState) ExecuteBefore() {
	if this.actor.GetProp().HpNow() <= 0 {
		this.fsm.Event(base.StateTriggerDeath)
	}
	this.actor.CastPassiveSkill()
}

func (this *FightState) Execute() {

	if this.inWalkTime() {
		return
	}

	var target base.Actor
	if this.context != nil {
		target = this.context.getFightStageTarget()
	} else {
		target = this.getFightStageTarget()
	}
	if target == nil {
		return
	}
	this.attackTarget(target)

}

func (this *FightState) attackTarget(target base.Actor) {
	dis := scene.DistanceByPoint(this.actor.Point(), target.Point())
	if dis <= 0 || !this.actor.Point().CanStand(this.actor) {
		if this.isCanMove(){
			monsterRandWalk(this.actor)
			this.fsm.SetSleepTime(400)
			return
		}
	}

	//判断技能是否可以释放
	skill := this.actor.GetReadySkill(target)
	if skill != nil {
		//logger.Debug("获取一个可释放的技能：%v,技能：%v,当前锁定目标：%v-%v", this.actor.GetObjId(), skill.Skillid, target.GetType(), target.NickName())
		dir := scene.GetFaceDirByPoint(this.actor.Point(), target.Point())
		this.actor.CastSkill(skill, dir, []int{target.GetObjId()}, false)
		this.fsm.SetSleepTime(500 + rand.Intn(200))
		return
	}

	if dis > 1 {
		this.chaseTarget(target)
	}
}

func (this *FightState) getFightStageTarget() base.Actor {

	//查找一个仇恨值最高的目标
	targetId := this.actor.GetTargetByFirstTheat()
	if targetId == 0 {
		//仇恨列表里面没有目标，重新进入活跃状态
		this.fsm.Event(base.StateTriggerActive)
		return nil
	}
	target := this.actor.GetFight().GetActorByObjId(targetId)
	if target == nil || target.GetProp().HpNow() <= 0 || !target.GetVisible() {
		//清除一个无效的仇恨目标，重新跟随新的目标
		this.actor.ClearTheatTarget(targetId)
		return nil
	}
	return target
}

//追击敌人
func (this *FightState) chaseTarget(target base.Actor) {
	//往目标移动
	if !this.isCanMove() {
		return
	}
	chaseDisT := 0
	chaseDis := 0
	if this.actor.GetType() == base.ActorTypeMonster {
		monsterActor := this.actor.(*actorPkg.MonsterActor)
		chaseDisT = monsterActor.MonsterT.ChaseDis
		chaseDis = scene.DistanceByPoint(monsterActor.BirthPoint(), target.Point())
	}
	if chaseDis > chaseDisT {
		this.actor.ClearTheatTarget(target.GetObjId())
		return
	}

	moveToPoint(this.actor, this.actor.Point(), target.Point(), constFight.MOVE_TYPE_CHASE)
	//TODO 优化 设置怪物移动间隔
	if this.actor.GetType() == pb.SCENEOBJTYPE_MONSTER {
		this.lastMoveTime = time.Now().UnixNano() + 1e9 + int64(rand.Intn(5e8))
	} else {
		this.lastMoveTime = time.Now().UnixNano()
	}
}

func (this *FightState) isCanMove() bool {

	if this.actor.GetType() == base.ActorTypeMonster {
		if this.actor.(*actorPkg.MonsterActor).MonsterT.Move != 1{
			return false
		}
	}
	if !this.actor.CanMove() {
		return false
	}
	return true
}

//判断是否到达可移动时间
func (this *FightState) inWalkTime() bool {
	timeStamp := time.Now().UnixNano()
	interval := GetActionInterval(300)
	if timeStamp-this.lastMoveTime >= interval {
		return false
	}
	//logger.Debug("timeStamp:%d, this.lastMoveTime:%d, interval:%d", timeStamp, this.lastMoveTime, interval)
	return true
}
