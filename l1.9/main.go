package main

import (
	"fmt"
	"sync"
)

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(firstChan)

		numbres := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, v := range numbres {
			firstChan <- v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(secondChan)

		for v := range firstChan {
			secondChan <- v * 2
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range secondChan {
			fmt.Printf("value: %v\n", v)
		}
		fmt.Println("Все данные обработаны")
	}()

	wg.Wait()
}
