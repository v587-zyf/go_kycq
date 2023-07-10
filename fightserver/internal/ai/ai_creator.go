package ai

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/fsm"
)

func NewUserAI(actor base.Actor) *fsm.FSM {
	fsmAI := fsm.NewFSM()
	activeState := NewPlayerAIActiveState(fsmAI, actor)
	fightState := NewFightState(fsmAI, actor, nil)
	dieState := NewPlayerDeathState(fsmAI, actor)
	//attackState := NewAttackState(fsmAI, actor)
	//deathState := NewDeathState(fsmAI, actor)
	fsmAI.SetInitState(activeState)
	fsmAI.Register(base.StateTriggerActive, []fsm.State{fightState,dieState}, activeState)
	fsmAI.Register(base.StateTriggerFight, []fsm.State{activeState}, fightState)
	//fsmAI.Register(StateTriggerAttack, []fsm.State{chaseState}, attackState)
	fsmAI.Register(base.StateTriggerDeath, []fsm.State{activeState, fightState}, dieState)
	return fsmAI
}

func NewUserFsm(actor base.Actor) *fsm.FSM {

	fsmAI := fsm.NewFSM()
	activeState := NewPlayerActiveState(fsmAI, actor)
	dieState := NewPlayerDeathState(fsmAI, actor)
	fsmAI.SetInitState(activeState)
	fsmAI.Register(base.StateTriggerActive, []fsm.State{dieState}, activeState)
	fsmAI.Register(base.StateTriggerDeath, []fsm.State{activeState}, dieState)
	return fsmAI
}

func NewMonsterAI(actor base.Actor) *fsm.FSM {
	fsmAI := fsm.NewFSM()
	activeState := NewMonsterActiveState(fsmAI, actor)
	fightState := NewMonsterFightState(fsmAI, actor)
	dieState := NewMonsterDieState(fsmAI, actor)
	fsmAI.SetInitState(activeState)
	fsmAI.Register(base.StateTriggerActive, []fsm.State{fightState, dieState}, activeState)
	fsmAI.Register(base.StateTriggerFight, []fsm.State{activeState}, fightState)
	//fsmAI.Register(StateTriggerAttack, []fsm.State{fightState}, attackState)
	fsmAI.Register(base.StateTriggerDeath, []fsm.State{activeState, fightState}, dieState)
	return fsmAI
}

//
func NewHeroAI(actor base.Actor) *fsm.FSM {
	fsmAI := fsm.NewFSM()
	//	followState := NewFollowState(fsmAI, actor)
	//	chaseState := NewChaseState(fsmAI, actor)
	//	attackState := NewAttackState(fsmAI, actor)
	//	deathState := NewDeathState(fsmAI, actor)
	//	fsmAI.SetInitState(followState)
	//	fsmAI.Register(StateTriggerActive, []fsm.State{followState}, chaseState)
	//	fsmAI.Register(StateTriggerActive, []fsm.State{chaseState}, followState)
	//	fsmAI.Register(StateTriggerAttack, []fsm.State{chaseState}, attackState)
	//	fsmAI.Register(StateTriggerDeath, []fsm.State{followState, chaseState, attackState}, deathState)
	return fsmAI
}
