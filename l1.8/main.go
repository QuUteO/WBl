package main

import "fmt"

func SetBit(num int64, position uint, bitValue int) int64 {
	if bitValue == 1 {
		return num | (1 << position)
	} else {
		return num &^ (1 << position)
	}
}

func main() {
	var num int64 = 5
	position := uint(0)

	fmt.Printf("num: %b\nnumResult: %b\n", num, SetBit(num, position, 0))
}
