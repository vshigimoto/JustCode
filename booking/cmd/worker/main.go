package main

import (
	"booking/internal/booking/worker"
)

func main() {
	go worker.Worker(worker.Jobs, worker.Result)
	time.Sle
}
