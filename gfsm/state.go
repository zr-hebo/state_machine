package gfsm

// Stater Stater
type Stater interface {
	Walk(input interface{}) (Stater, interface{}, error)
	String() string
	IsEnd() bool
	GetVal() interface{}
	SetMachine(*StateMachine)
}

// SimpleState State
type SimpleState struct {
	name string
	sm   *StateMachine
}

// Walk Walk
func (s *SimpleState) Walk(val interface{}) (
	nextState *SimpleState, option interface{}, err error) {
	return
}

func (s *SimpleState) String() (str string) {
	str = s.name
	return
}

// GetVal GetVal
func (s *SimpleState) GetVal() interface{} {
	return s.name
}

// SetMachine SetMachine
func (s *SimpleState) SetMachine(sm *StateMachine) {
	s.sm = sm
}
