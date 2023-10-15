package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	// Горутина 1 захватывает мьютекс
	go func() {
		fmt.Println("Goroutine 1: Trying to acquire the lock")
		mu.Lock()
		fmt.Println("Goroutine 1: Lock acquired")
	}()

	// Горутина 2 также пытается захватить мьютекс
	go func() {
		fmt.Println("Goroutine 2: Trying to acquire the lock")
		mu.Lock()
		fmt.Println("Goroutine 2: Lock acquired")
	}()

	// Ожидаем завершения горутин
	select {}
}
