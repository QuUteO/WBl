package main

import (
	"fmt"
	"time"
)

func channels(id int, ch <-chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Channel", id)
			return
		default:
			id++
			fmt.Println("Channel", id)
		}
	}
}

func main() {
	stopChan := make(chan struct{})
	id := 0
	go channels(id, stopChan)

	time.Sleep(1 * time.Second)
	close(stopChan) // или stopChan <- struct{}{}
	time.Sleep(1 * time.Second)
}
