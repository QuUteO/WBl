package main

import "fmt"

func main() {
	a, b := 5, 10

	a = a + b // a = 15, b = 10
	b = a - b // a = 15, b = 5
	a = a - b // a = 10, b = 5

	fmt.Println(a, b)
}
