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
	sm.currentState = ss
	// sm.startState = ss
	// sm.endState = es
	sm.allStates = allStates
	return
}

// Walk Walk
func (sm *StateMachine) Walk(input interface{}) (err error) {
	if sm.currentState == nil {
		err = errors.New("state machine's state is NULL, cannot walk to")
		return
	}

	if sm.IsFinished() {
		err = errors.New(
			"state machine was already in finish state, cannot walk")
		return
	}

	nextState, err := sm.currentState.Walk(input)
	if err != nil {
		return
	}

	sm.currentState = nextState
	return
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
