package gfsm

import (
	"fmt"
	"testing"
)

type originState struct {
	SimpleState
}

type closeDoorState struct {
	SimpleState
}

type openWindowState struct {
	SimpleState
}

func (os *originState) Walk(val ...interface{}) (
	nextState Stater, err error) {
	fmt.Println(os.name, ":in origin state")
	fmt.Println(val)
	cds := &closeDoorState{}
	cds.name = "open"
	nextState = cds
	return
}

func (cds *closeDoorState) Walk(val ...interface{}) (
	nextState Stater, err error) {
	fmt.Println(cds.name, ":in close door state")
	fmt.Println(val)
	ows := &openWindowState{}
	ows.name = "safe"
	nextState = ows
	return
}

func (ows *openWindowState) Walk(val ...interface{}) (
	nextState Stater, err error) {
	fmt.Println(val)
	err = fmt.Errorf("the end state cannot walk any more")
	return
}

func (ows *openWindowState) IsEnd() bool {
	return true
}

func Test_StateMachine(t *testing.T) {
	startState := &originState{}
	startState.name = "init"
	sm := NewStateMachine(startState)
	for !sm.IsFinished() {
		err := sm.Walk("haha")
		if err != nil {
			t.Fatal(err)
		}
	}
}
