package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	m  map[string]int
	mu sync.Mutex
}

func NewMyMap() *MyMap {
	return &MyMap{
		m: make(map[string]int),
	}
}

func (m *MyMap) Set(key string, value int) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

func (m *MyMap) Get(key string) int {
	m.mu.Lock()
	value := m.m[key]
	m.mu.Unlock()
	return value
}

func main() {
	m := NewMyMap()
	m.Set("key1", 100)
	m.Set("key2", 200)
	fmt.Println(m.Get("key1")) // 100
	fmt.Println(m.Get("key2")) // 200
}
