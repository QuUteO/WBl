package main

import (
	"fmt"
	"runtime"
	"time"
)

func RunTime(id int) {
	defer fmt.Println("GoRoutine is Exit")

	for i := 0; i < id; i++ {
		time.Sleep(1 * time.Second)
		if i == 2 {
			fmt.Printf("Worker %d calling Goexit()\n", id)
			runtime.Goexit()
		}
	}
}

func main() {
	go RunTime(5)
	time.Sleep(5 * time.Second)
}
