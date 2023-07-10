package ai

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/fsm"
	"time"
)

type PlayerAIActiveState struct {
	*ActivieState
	randMoveTime int64
}

func NewPlayerAIActiveState(aifsm *fsm.FSM, actor base.Actor) *PlayerAIActiveState {
	return &PlayerAIActiveState{
		ActivieState: &ActivieState{
			actor:    actor,
			fsm:      aifsm,
			findTime: time.Now().Unix(),
		},
	}
}

func (this *PlayerAIActiveState) Enter() {

	this.randMoveTime = 0
	//清空路径
	this.actor.SetPathSlice(nil)
}

func (this *PlayerAIActiveState) Leave() {
	//清空路径
	this.actor.SetPathSlice(nil)
}

func (this *PlayerAIActiveState) Execute() {

	this.ActivieState.Execute()

	allObjs := this.actor.GetScene().GetSceneAllObj()
	minDis := 99999999
	var enemy base.Actor
	for _, v := range allObjs {
		if v.IsSceneObj() {
			continue
		}

		if findActor, ok := v.GetContext().(base.Actor); ok {

			if findActor.GetProp().HpNow() > 0 && findActor.CanAttack() && findActor.GetVisible() && this.actor.IsEnemy(findActor) {
				dis := scene.DistanceByPoint(this.actor.Point(), v.Point())
				if minDis > dis {
					minDis = dis
					enemy = findActor
				}
			}
		}
	}

	if enemy != nil {
		this.actor.AddTheat(enemy.GetObjId(), 0)
		this.fsm.Event(base.StateTriggerFight)
	}
	this.fsm.SetSleepTime(400)
}
