package main

import (
	"fmt"
	"strings"
)

func allUnique(s string) bool {
	seen := make(map[rune]bool)

	// приводим строку к одному регистру (нижний/верхний)
	s = strings.ToLower(s)

	for _, r := range s {
		if seen[r] {
			return false // символ уже встречался
		}
		seen[r] = true
	}
	return true
}

func main() {
	fmt.Println(allUnique("abcd"))      // true
	fmt.Println(allUnique("abCdefAaf")) // false
	fmt.Println(allUnique("aabcd"))     // false
}
