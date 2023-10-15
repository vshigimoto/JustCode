package main

import (
	"sync"
)

func mergeChannels(channels []chan int) <-chan int {
	mergedChannel := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				mergedChannel <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChannel)
	}()

	return mergedChannel
}

func main() {
	// Пример использования
	n := 5 // Количество каналов
	channels := make([]chan int, n)

	for i := 0; i < n; i++ {
		channels[i] = make(chan int)
		go func(ch chan int, value int) {
			defer close(ch)
			ch <- value
		}(channels[i], i)
	}

	mergedChannel := mergeChannels(channels)

	for value := range mergedChannel {
		// Обработка данных из объединенного канала
		// В этом примере просто выводим значения на экран
		println(value)
	}
}
