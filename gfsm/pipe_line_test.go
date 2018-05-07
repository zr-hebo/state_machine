package gfsm

import (
	"context"
	"fmt"
	"testing"
)

func Test_ExecuteFullPipeLine(t *testing.T) {
	task := &testTask{}
	stepLogger := &testStepLogger{}
	pipeLine := NewPipeLine(task, stepLogger)
	pipeLine.RegistInitStep(func() error {
		fmt.Println("xiaoming in home")
		return nil
	})
	pipeLine.RegistFinallyStep(func() error {
		fmt.Println("xiaoming sit down")
		return nil
	})

	pipeLine.RegistStep("close door", func() error {
		fmt.Println("xiaoming close door")
		return nil
	})
	pipeLine.RegistStep("open window", func() error {
		fmt.Println("xiaoming open window")
		return nil
	})
	ctx := context.Background()
	err := pipeLine.Execute(ctx)
	if err != nil {
		t.Log(err)
	}
}

func Test_ExecuteNoInitStepPipeLine(t *testing.T) {
	task := &testTask{}
	stepLogger := &testStepLogger{}
	pipeLine := NewPipeLine(task, stepLogger)

	pipeLine.RegistFinallyStep(func() error {
		fmt.Println("xiaoming sit down")
		return nil
	})

	pipeLine.RegistStep("close door", func() error {
		fmt.Println("xiaoming close door")
		return nil
	})
	pipeLine.RegistStep("open window", func() error {
		fmt.Println("xiaoming open window")
		return nil
	})
	ctx := context.Background()
	err := pipeLine.Execute(ctx)
	if err != nil {
		t.Log(err)
	}
}

func Test_ExecuteNoFinallyPipeLine(t *testing.T) {
	task := &testTask{}
	stepLogger := &testStepLogger{}
	pipeLine := NewPipeLine(task, stepLogger)
	pipeLine.RegistInitStep(func() error {
		fmt.Println("xiaoming in home")
		return nil
	})
	pipeLine.RegistStep("close door", func() error {
		fmt.Println("xiaoming close door")
		return nil
	})
	pipeLine.RegistStep("open window", func() error {
		fmt.Println("xiaoming open window")
		return nil
	})
	ctx := context.Background()
	err := pipeLine.Execute(ctx)
	if err != nil {
		t.Log(err)
	}
}

type testTask struct {
}

func (tt *testTask) GetStep() (int, error) {
	return 0, nil
}

func (tt *testTask) UpdateStatus(status int) error {
	return nil
}

func (tt *testTask) UpdateTaskInfo(step, status int) error {
	return nil
}

type testStepLogger struct {
}

func (tsl *testStepLogger) Log(
	stepOrder int, stepName string, status int, info string) (err error) {
	fmt.Printf("Log step %d. %s: execute status:%d, message:%s\n",
		stepOrder, stepName, status, info)
	return nil
}
