package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: channel closed, exiting\n", id)
				return
			}
			fmt.Printf("Worker %d job %d\n", id, job)
			time.Sleep(500 * time.Millisecond) // Имитация работы
			fmt.Printf("Worker %d  job %d\n", id, job)

		case <-ctx.Done():
			fmt.Printf("Worker %d received shutdown signal\n", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go <num_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	// Запуск workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// Генерация jobs
	go func() {
		jobID := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				jobs <- jobID
				fmt.Printf("Sent job %d to queue\n", jobID)
				jobID++
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// Ожидание сигнала
	<-sigChan
	fmt.Println("\nShutting down gracefully...")

	// Отмена контекста и закрытие канала
	cancel()
	close(jobs)

	// Ожидание завершения workers
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Таймаут для graceful shutdown
	select {
	case <-done:
		fmt.Println("All workers completed successfully")
	case <-time.After(5 * time.Second):
		fmt.Println("Shutdown timeout, forcing exit")
	}
}
