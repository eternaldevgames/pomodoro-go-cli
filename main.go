package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var currentTask Pomodoro = Pomodoro{}

func main() {

	taskPtr := flag.String("task", "", "Task to add. (Required)")
	flag.Parse()

	if *taskPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	currentTask = AddTask(*taskPtr)
	fmt.Println(PrintTask(currentTask))

	fmt.Printf("Start the task! Focus on %s\n", *taskPtr)

	timer1 := time.NewTimer(25 * time.Second)
	<-timer1.C

	fmt.Println("Congrats! Task time is complete. Take a break")
}
