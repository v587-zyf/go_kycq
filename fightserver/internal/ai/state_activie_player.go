package ai

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/fsm"
	"time"
)

type PlayerActiveState struct {
	*ActivieState
}

func NewPlayerActiveState(aifsm *fsm.FSM, actor base.Actor) *PlayerActiveState {
	return &PlayerActiveState{
		ActivieState: &ActivieState{
			actor:    actor,
			fsm:      aifsm,
			findTime: time.Now().Unix(),
		},
	}
}

func (this *PlayerActiveState) Execute() {

	this.actor.CastPassiveSkill()
}
