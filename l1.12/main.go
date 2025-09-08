package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]bool)
	for _, v := range sequence {
		m[v] = true
	}

	var result []string
	for k := range m {
		result = append(result, k)
	}

	fmt.Println(result)
}
