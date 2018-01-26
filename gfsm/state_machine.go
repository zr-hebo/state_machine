package gfsm

import (
	"errors"
	"fmt"
)

// StateMachine StateMachine
type StateMachine struct {
	currentState Stater
}

// NewStateMachine NewStateMachine
func NewStateMachine(startState Stater) (sm *StateMachine) {
	sm = new(StateMachine)
	startState.SetMachine(sm)
	sm.currentState = startState
	return
}

// Walk Walk
func (sm *StateMachine) Walk(input interface{}) (option interface{}, err error) {
	if sm.currentState == nil {
		err = errors.New("state machine's state is NULL, cannot walk to")
		return
	}

	nextState, option, err := sm.currentState.Walk(input)
	if err != nil {
		return
	}

	sm.currentState = nextState
	return
}

// GetState GetState
func (sm *StateMachine) GetState() Stater {
	return sm.currentState
}

// SetState SetState
/* func (sm *StateMachine) SetState(state Stater) {
	sm.currentState = state
} */

func (sm *StateMachine) String() (str string) {
	str = fmt.Sprintf("current state is %s", sm.currentState)
	return
}

// IsFinished IsFinished
func (sm *StateMachine) IsFinished() bool {
	if sm.currentState == nil {
		return true
	}

	return sm.currentState.IsEnd()
}
