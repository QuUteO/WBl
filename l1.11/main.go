package main

import "fmt"

func intersection(a []int, b []int) []int {
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}

	var result []int
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	result := intersection(a, b)

	fmt.Println(result)
}
