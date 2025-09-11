package main

import (
	"fmt"
	"math/big"
)

// Add функция сложения math/big
func Add(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}

// Substraction функция вычитания math/big
func Substraction(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Sub(a, b)
	return result
}

// Divide фукнция деления
func Divide(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Div(a, b)
	return result
}

// Multiply функция умножения
func Multiply(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(a, b)
	return result
}

func main() {
	a, b := big.NewInt(2), big.NewInt(2)

	a.Exp(a, big.NewInt(23), nil) // a == 2^22
	b.Exp(b, big.NewInt(21), nil) // b == 2^21

	// сложение
	sum := Add(a, b)
	// вычитание
	sub := Substraction(a, b)
	// деление
	div := Divide(sub, b)
	// умножение
	mul := Multiply(div, b)

	// Выводим результаты
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Сложение: a + b = %s\n", sum.String())
	fmt.Printf("Вычитание: a - b = %s\n", sub.String())
	fmt.Printf("Умножение: a * b = %s\n", mul.String())
	fmt.Printf("Деление: a / b = %s\n", div.String())

}
