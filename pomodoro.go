package main

import (
	"time"
)

type Pomodoro struct {
	// Name of the task
	TaskName string

	// StartTime of the task created by the user
	StartTime time.Time
}

func AddTask(taskName string) Pomodoro {

	return Pomodoro{
		TaskName:  taskName,
		StartTime: time.Now(),
	}

}

func PrintTask(task Pomodoro) string {
	// Check for empty struct
	if (Pomodoro{} == task) {
		return "Current task is empty"
	} else {
		return "Task - " + task.TaskName
	}
}
