package main

import "testing"

func TestAddTask(t *testing.T) {
	expectedTaskName := "Testing task"
	newTask := AddTask(expectedTaskName)

	if newTask.TaskName != expectedTaskName {
		t.Errorf("TestAddTask Failed - Expected: %s Got: %s", expectedTaskName, newTask.TaskName)
	}
}

func TestPrintTask_WithTask(t *testing.T) {
	expectedTaskName := "Testing Task"
	newTask := AddTask("Testing Task")

	printOutput := PrintTask(newTask)
	expectedPrintOutput := "Task - " + expectedTaskName

	if printOutput != expectedPrintOutput {
		t.Errorf("PrintTask Failed - Expected: %s Got: %s", expectedPrintOutput, printOutput)
	}
}

func TestPrintTask_Empty(t *testing.T) {

	printOutput := PrintTask(Pomodoro{})
	expectedPrintOutput := "Current task is empty"

	if printOutput != expectedPrintOutput {
		t.Errorf("PrintTask Failed - Expected: %s Got: %s", expectedPrintOutput, printOutput)
	}
}
