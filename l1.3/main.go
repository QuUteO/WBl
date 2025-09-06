package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan int) {
	for n := range jobs {
		fmt.Printf("worker %d received job %d\n", id, n)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go <file>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid argument")
		return
	}

	jobs := make(chan int, 100)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs)
	}

	jobID := 1
	for {
		jobs <- jobID
		jobID++
		time.Sleep(200 * time.Millisecond)
	}

}
