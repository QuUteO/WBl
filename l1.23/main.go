package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50}
	i := 2 // удалим элемент с индексом 2 (т.е. "30")

	// Сдвигаем хвост влево на одну позицию
	copy(slice[i:], slice[i+1:])

	// Обрезаем последний элемент
	slice = slice[:len(slice)-1]

	fmt.Println(slice) // [10 20 40 50]
}
