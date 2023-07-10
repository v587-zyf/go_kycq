package ai

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/fsm"
)

type DieState struct {
	fsm.DefaultState
	actor base.Actor
	fsm   *fsm.FSM
	isDie bool
}

func NewDeathState(aifsm *fsm.FSM, actor base.Actor) *DieState {
	return &DieState{
		actor: actor,
		fsm:   aifsm,
		isDie: false,
	}
}

func (this *DieState) Enter() {
	this.fsm.Pause()
}
