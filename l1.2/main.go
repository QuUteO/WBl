package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 5, 6, 8, 10}
	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			square := n * n
			fmt.Printf("%v: square: %v\n", n, square)
		}(number)
	}

	wg.Wait()
}
