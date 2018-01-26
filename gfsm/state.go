package gfsm

// Stater Stater
type Stater interface {
	Walk(input interface{}) (Stater, error)
	String() string
	IsEnd() bool
	SetOwnMachine(*StateMachine)
}

// SimpleState State
type SimpleState struct {
	name string
	sm   *StateMachine
}

// Walk Walk
func (s *SimpleState) Walk(val interface{}) (nextStat *SimpleState, err error) {
	return
}

func (s *SimpleState) String() (str string) {
	str = s.name
	return
}

// GetState GetState
func (s *SimpleState) GetState() (str string) {
	str = s.name
	return
}

// SetOwnMachine SetOwnMachine
func (s *SimpleState) SetOwnMachine(sm *StateMachine) {
	s.sm = sm
}
