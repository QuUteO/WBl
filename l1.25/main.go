package main

import (
	"fmt"
	"time"
)

func SleepChannel(d time.Duration) {
	ch := make(chan struct{})
	go func() {
		timer := time.NewTimer(d)
		<-timer.C
		ch <- struct{}{}
	}()
	// блокируемся, потому что ждем пока в канал придет сигнал
	<-ch
}

func SleepCicle(d time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= d {
			return
		}
	}
}

func main() {
	fmt.Println("Start:", time.Now())
	SleepChannel(2 * time.Second)
	fmt.Println("End:", time.Now())

	fmt.Println("Start:", time.Now())
	SleepCicle(2 * time.Second)
	fmt.Println("End:", time.Now())
}
