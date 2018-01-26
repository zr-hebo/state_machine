package gfsm

// Stater Stater
type Stater interface {
	Walk(input interface{}) (Stater, error)
	String() string
	IsEnd() bool
	SetMachine(*StateMachine)
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

// SetMachine SetMachine
func (s *SimpleState) SetMachine(sm *StateMachine) {
	s.sm = sm
}
