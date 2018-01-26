package gfsm

import (
	"errors"
	"fmt"
)

// StateMachine StateMachine
type StateMachine struct {
	currentState Stater
	// startState   Stater
	// endState     Stater
	allStates []Stater
}

// NewStateMachine NewStateMachine
func NewStateMachine(ss Stater, allStates []Stater) (sm *StateMachine) {
	sm = new(StateMachine)

	for _, s := range allStates {
		s.SetMachine(sm)
	}

	sm.currentState = ss
	sm.allStates = allStates
	return
}

// Walk Walk
func (sm *StateMachine) Walk(input interface{}) (err error) {
	if sm.currentState == nil {
		err = errors.New("state machine's state is NULL, cannot walk to")
		return
	}

	nextState, err := sm.currentState.Walk(input)
	if err != nil {
		return
	}

	sm.currentState = nextState
	return
}

// SetState SetState
func (sm *StateMachine) SetState(state Stater) {
	sm.currentState = state
}

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
