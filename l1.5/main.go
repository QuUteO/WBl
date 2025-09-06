package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// чтение из аргумента терминала
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go <file>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Invalid argument")
		return
	}

	// создаем канал для чтения и записи
	ch := make(chan int)

	// горутина для чтения
	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	// каждые 500 милисекунд будет отправляться в канал данные
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// время через которое завершится программа
	timeout := time.After(time.Duration(n) * time.Second)

	i := 1
	for {
		select {
		case <-ticker.C:
			ch <- i
			i++
		case <-timeout:
			close(ch) // закрываем канал чтоб не было чтения
			fmt.Println("Timeout")
			return
		}
	}

}
