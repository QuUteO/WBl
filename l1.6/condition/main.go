package main

import (
	"fmt"
	"time"
)

func condition(id int, stop *bool) {
	for *stop {
		id++
		fmt.Printf("Condition working %v %v\n", id, *stop)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Condition stopped")
}

func main() {
	stop := true
	id := 0
	go condition(id, &stop)

	time.Sleep(1 * time.Second)
	stop = false
	time.Sleep(1 * time.Second)
}
