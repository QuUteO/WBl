package main

import (
	"fmt"
	"strconv"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

// Save запись
func (m *SafeMap) Save(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = value
}

// Load чтение
func (m *SafeMap) Load(key string) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.m[key]
	if !ok {
		return 0
	}
	return v
}

// Len длина map
func (m *SafeMap) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.m)
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	// конкурентная запись
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			key := strconv.Itoa(id)
			safeMap.Save(key, id)
			wg.Done()
		}(i)
	}

	// конкурентное чтение
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			key := strconv.Itoa(id)
			val := safeMap.Load(key)
			fmt.Println(val)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("len mutex map %v\n", safeMap.Len())
}
