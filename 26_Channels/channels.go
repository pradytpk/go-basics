package main

import (
	"fmt"
	"time"
)

type ControlMessage int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMessage) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("Unhandled Control Message")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}
func main() {
	//jobs
	jobs := make(chan Job, 50)
	//results
	results := make(chan Result, 50)
	//controls
	controls := make(chan ControlMessage)

	go doubler(jobs, results, controls)
	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}
	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timed out")
			controls <- DoExit
			<-controls
			fmt.Println("Program exit")
			return
		}
	}
}
