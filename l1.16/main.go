package main

import "fmt"

// сортировка вставками
func quickSort(array []int) []int {
	length := len(array)

	for i := 1; i < length; i++ {
		j := i

		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}

	return array
}

func main() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42, 33, 10, 5, 99}

	res := quickSort(numbers)
	fmt.Println(res)
}
