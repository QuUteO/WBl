package main

import (
	"fmt"
)

// reverseWordsInPlace переворачивает слова на месте (в byte slice)
func reverseWordsInPlace(s string) string {
	// Конвертируем строку в slice байт для модификации
	bytes := []byte(s)

	// Сначала переворачиваем всю строку
	reverseBytes(bytes, 0, len(bytes)-1)

	// Затем переворачиваем каждое слово отдельно
	start := 0
	for i := 0; i <= len(bytes); i++ {
		if i == len(bytes) || bytes[i] == ' ' {
			reverseBytes(bytes, start, i-1)
			start = i + 1
		}
	}

	return string(bytes)
}

// reverseBytes переворачивает часть slice байт
func reverseBytes(bytes []byte, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func main() {
	input := "snow dog sun"
	output := reverseWordsInPlace(input)

	fmt.Printf("Вход:  \"%s\"\n", input)
	fmt.Printf("Выход: \"%s\"\n", output)

	// Дополнительные тесты
	examples := []string{
		"hello world",
		"the quick brown fox",
		"Go programming language",
	}

	for _, ex := range examples {
		result := reverseWordsInPlace(ex)
		fmt.Printf("\n«%s» → «%s»", ex, result)
	}
}
