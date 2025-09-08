package main

import (
	"fmt"
	"strings"
)

// Имитация функции создания большой строки
func createHugeString(size int) string {
	return strings.Repeat("x", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // 1024 байта

	// Корректное взятие подстроки без утечки памяти
	justString = string([]byte(v[:100]))

	// Альтернатива:
	// justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Printf("Длина justString: %d\n", len(justString))
	fmt.Printf("Первые 10 символов: %s\n", justString[:10])
}
