package ai

import (
	"time"

	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/fsm"
)

type ActivieState struct {
	fsm.DefaultState
	actor    base.Actor
	fsm      *fsm.FSM
	findTime int64
}

func NewActiveState(aifsm *fsm.FSM, actor base.Actor) *ActivieState {
	return &ActivieState{
		actor:    actor,
		fsm:      aifsm,
		findTime: time.Now().Unix(),
	}
}

func (this *ActivieState) ExecuteBefore() {

	if this.actor.GetProp().HpNow() <= 0 {
		this.fsm.Event(base.StateTriggerDeath)
	}
}
