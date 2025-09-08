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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	id := 0
	go contx(id, ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
