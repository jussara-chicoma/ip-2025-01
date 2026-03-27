package main

import (
	"fmt"
)

func main() {
	var a1, r, n int
	var soma int

	// Leitura dos valores
	fmt.Scan(&a1, &r, &n)

	// Soma dos n termos usando repetição
	for i := 0; i < n; i++ {
		termo := a1 + i*r
		soma += termo
	}

	// Saída
	fmt.Println(soma)
}
