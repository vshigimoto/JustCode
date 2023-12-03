package worker

import (
	"time"
)

var Jobs = make(chan int, 5)
var Result = make(chan int, 5)

func Worker(jobs <-chan int, result chan<- int) {
	timeout := time.After(5 * time.Minute)
	for {
		select {
		case <-jobs:
			res := <-jobs
			result <- res
		case <-timeout:
			return
		}
	}
}
