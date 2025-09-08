package main

import "fmt"

func detectTypeAdvanced(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	}
	return "unknown"
}

func main() {
	values := []interface{}{
		42,
		"hello",
		true,
		make(chan int),
		make(chan string),
		make(chan bool),
		nil,
	}

	for _, v := range values {
		res := detectTypeAdvanced(v)
		fmt.Println(res)
	}
}
