package main

import (
	"context"
	"fmt"
	"time"
)

func contx(id int, context context.Context) {
	for {
		select {
		case <-context.Done():
			fmt.Println("Context done")
			return
		default:
			id++
			fmt.Println("Context", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	id := 0
	go contx(id, ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}
