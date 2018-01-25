package gfsm

// Stater Stater
type Stater interface {
	Walk(input interface{}) (Stater, error)
	String() string
	IsEnd() bool
}

// SimpleState State
type SimpleState struct {
	name string
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
