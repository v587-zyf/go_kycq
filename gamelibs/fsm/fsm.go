package fsm

import (
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"errors"
)

type State interface {
	Enter()
	ExecuteBefore()
	Execute()
	Leave()
}

type DefaultState struct {
}

type Trigger int

type FSM struct {
	curState    State
	transitions map[Trigger]map[State]State
	isRun       bool
	sleepTime   int64
}

func NewFSM() *FSM {
	return &FSM{
		transitions: make(map[Trigger]map[State]State),
		isRun:       false,
	}
}

/**
 *  @Description: 设置休眠时间
 *  @param sleeptime 毫秒
 */
func (this *FSM) SetSleepTime(sleeptime int) {

	this.sleepTime = common.GetNowMillisecond() + int64(sleeptime)
}

func (this *FSM) Start(initState State) {
	this.curState = initState
	initState.Enter()
	initState.Execute()
}

func (this *FSM) Run() {
	if this.curState != nil && this.isRun {
		if common.GetNowMillisecond() > this.sleepTime {
			this.curState.ExecuteBefore()
			this.curState.Execute()
		}
	}
}

func (this *FSM) Stop() {
	if this.curState != nil {
		this.curState.Leave()
	}
}

func (this *FSM) Pause() {
	this.isRun = false
}

func (this *FSM) Recover() {
	this.isRun = true
}

func (this *FSM) GetRunState() bool {
	return this.isRun
}

func (this *DefaultState) Enter() {
}

func (this *DefaultState) ExecuteBefore() {

}

func (this *DefaultState) Execute() {
}

func (this *DefaultState) Leave() {
}

func (this *FSM) Register(trigger Trigger, fromStates []State, toState State) {
	var stateMap map[State]State
	var ok bool
	for _, fromState := range fromStates {
		if stateMap, ok = this.transitions[trigger]; !ok {
			stateMap = make(map[State]State)
			this.transitions[trigger] = stateMap
		}
		stateMap[fromState] = toState
	}
}

func (this *FSM) SetInitState(state State) {
	this.curState = state
}

func (this *FSM) Event(trigger Trigger) error {
	if stateMap, ok := this.transitions[trigger]; ok {
		if state, ok := stateMap[this.curState]; ok {
			this.curState.Leave()
			this.curState = state
			state.Enter()
			return nil
		}
	}
	return errors.New("no state switch found")
}

func (this *FSM) Info(){
	logger.Debug("<------------------------->",this.isRun,this.curState != nil,common.GetNowMillisecond() > this.sleepTime)
}
