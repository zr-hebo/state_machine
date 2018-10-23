package gfsm

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
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
	task        taskTracker
	initStep    *Step
	finallyStep *Step
	logger      stepLogger
	steps       []*Step
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

// RegistInitStep regist initial step to pipe line
func (pl *PipeLine) RegistInitStep(initFunc func() error) {
	pl.initStep = NewStep("initial execute context", initFunc)
}

// RegistFinallyStep regist finally step to pipe line
func (pl *PipeLine) RegistFinallyStep(finallyFunc func() error) {
	pl.finallyStep = NewStep("execute finally step", finallyFunc)
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
func (pl *PipeLine) Execute(ctx context.Context) (err error) {
	beginStep, err := pl.task.GetStep()
	if err != nil {
		return
	}

	// steps := make([]*Step, 0, len(pl.steps)+1)
	// steps = append(steps, pl.initStep)
	// steps = append(steps, pl.steps...)
	steps := pl.steps

	// 结束时，执行清理步骤
	defer func() {
		if pl.finallyStep == nil {
			return
		}

		finallyErr := execStep(pl.finallyStep)
		pl.logStep(-1, pl.finallyStep.name, finallyErr)
	}()

	pl.task.UpdateStatus(StatusRunning)
	// 执行初始化步骤
	if pl.initStep != nil {
		err = execStep(pl.initStep)
		pl.logStep(0, pl.initStep.name, err)
		if err != nil {
			pl.task.UpdateStatus(StatusFailed)
			return
		}
	}

	for idx, step := range steps {
		select {
		case <-ctx.Done():
			err = errors.New("context cancel pipe line execute")
			return
		default:
		}

		if (idx < beginStep) || step == nil {
			continue
		}

		if err = execStep(step); err != nil {
			pl.task.UpdateTaskInfo(idx, StatusFailed)
			pl.logStep(idx, step.name, err)
			return
		}

		pl.task.UpdateTaskInfo(idx+1, StatusRunning)
		pl.logStep(idx, step.name, nil)
	}

	pl.task.UpdateStatus(StatusOK)
	return
}

func (pl PipeLine) logStep(
	stepOrder int, stepName string, err error) {
	if pl.logger == nil {
		return
	}

	status := StatusOK
	info := InfoOK
	if err != nil {
		status = StatusFailed
		info = err.Error()
	}

	_ = pl.logger.Log(stepOrder, stepName, status, info)
}

func (pl PipeLine) String() (str string) {
	steps := make([]string, 0, len(pl.steps))
	for idx, step := range pl.steps {
		steps = append(steps, fmt.Sprintf("Step%d: %s", idx+1, step.Name()))
	}
	return strings.Join(steps, "; ")
}

func execStep(step *Step) (err error) {
	defer func() {
		if reason := recover(); reason != nil {
			trackBack := string(debug.Stack())
			err = fmt.Errorf("panic in step of pipe line, %v", trackBack)
		}
	}()

	if step == nil {
		return
	}

	err = step.execFunc()
	return
}
