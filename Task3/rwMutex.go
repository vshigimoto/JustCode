package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем RWMutex
	rwm := sync.RWMutex{}

	// Создаем счетчик
	counter := 0

	// Создаем два потока
	go func() {
		// Читаем счетчик
		rwm.RLock()
		fmt.Println(counter)
		rwm.RUnlock()
	}()
	go func() {
		// Изменяем счетчик
		rwm.Lock()
		counter++
		rwm.Unlock()
	}()

	// Ждем завершения потоков
	time.Sleep(1 * time.Second)

	// Получаем значение счетчика
	fmt.Println(counter) // 1
}
