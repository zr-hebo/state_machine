package gfsm

import (
	"fmt"
	"strings"
)

const (
	// InfoOK ok message
	InfoOK = "OK"
	// StatusFailed task status failed
	StatusFailed = -1
	// StatusReady task status ready
	StatusReady = 0
	// StatusRunning task status running
	StatusRunning = 1
	// StatusOK task status ok
	StatusOK = 2
)

// taskTracker keep track of executing of task
type taskTracker interface {
	GetStep() (int, error)
	UpdateStatus(status int) error
	UpdateTaskInfo(step, status int) error
}

// Step execute step in task
type Step struct {
	name     string
	execFunc func() error
}

// StepLogger
type stepLogger interface {
	Log(stepOrder int, stepName string, status int, info string) error
}

// Name get step name
func (s *Step) Name() string {
	return s.name
}

// PipeLine get step name
type PipeLine struct {
	task   taskTracker
	logger stepLogger
	steps  []*Step
}

// NewStep create step in task
func NewStep(name string, execFunc func() error) (step *Step) {
	step = &Step{
		name:     name,
		execFunc: execFunc,
	}
	return
}

// NewPipeLine create new piple line
func NewPipeLine(task taskTracker, logger stepLogger) (pl *PipeLine) {
	pl = new(PipeLine)
	pl.steps = make([]*Step, 0, 16)
	pl.task = task
	pl.logger = logger
	return
}

// RegistStep regist step name and exec to pipe line
func (pl *PipeLine) RegistStep(stepName string, execFunc func() error) {
	step := NewStep(stepName, execFunc)
	pl.steps = append(pl.steps, step)
}

// AddStep add step to pip line
func (pl *PipeLine) AddStep(step *Step) {
	pl.steps = append(pl.steps, step)
}

// Execute execute each step in pipe line
func (pl *PipeLine) Execute() (err error) {
	beginStep, err := pl.task.GetStep()
	if err != nil {
		return
	}

	for idx, step := range pl.steps {
		if idx < beginStep {
			continue
		}

		err = step.execFunc()
		if err != nil {
			pl.task.UpdateTaskInfo(idx, StatusFailed)
			pl.logger.Log(idx, step.name, StatusFailed, err.Error())
			return
		}

		pl.task.UpdateTaskInfo(idx+1, StatusRunning)
		pl.logger.Log(idx, step.name, StatusOK, InfoOK)
	}

	pl.task.UpdateStatus(StatusOK)
	return
}

func (pl PipeLine) String() (str string) {
	steps := make([]string, 0, len(pl.steps))
	for idx, step := range pl.steps {
		steps = append(steps, fmt.Sprintf("Step%d: %s", idx, step.Name()))
	}
	return strings.Join(steps, "; ")
}
