package main

import "fmt"

func binarySearch(array []int, target int) int {
	return binarySearchRecursive(array, target, 0, len(array)-1)
}

func binarySearchRecursive(array []int, target int, left, right int) int {
	// если элемент не найден
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	// если элемент в середине находится
	if array[mid] == target {
		return mid
	}

	if array[mid] > target {
		return binarySearchRecursive(array, target, left, mid-1)
	} else {
		return binarySearchRecursive(array, target, mid+1, right)
	}
}

func main() {
	array := []int{5, 10, 11, 12, 22, 25, 33, 34, 42, 50, 64, 76, 88, 90, 99}
	target := 90

	fmt.Println(binarySearch(array, target))
}
