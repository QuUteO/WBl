package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewConstructor(value int) *Counter {
	return &Counter{
		value: value,
	}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += 1
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := NewConstructor(1)
	var wg sync.WaitGroup

	NewGoroutines := 10
	wg.Add(NewGoroutines)

	for i := 0; i < NewGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			// каждая горутина инкрементирует 10 раз
			for j := 0; j < 10; j++ {
				counter.Increment()
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("Итоговое значение счётчика: %d\n", counter.Value())
	fmt.Printf("Ожидаемое значение: %v\n", NewGoroutines*10+1)
}
