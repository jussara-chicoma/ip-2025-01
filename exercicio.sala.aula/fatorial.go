package main

import (
	"fmt"
)

// Fatorial calcula o fatorial de um número inteiro não negativo
func Fatorial(n int) int {
	if n < 0 {
		panic("Número negativo não é permitido")
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Fatorial(n-1)
}

func main() {
	fmt.Println(Fatorial(5)) // Saída: 120
}
