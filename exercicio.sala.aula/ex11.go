package main

import (
	"fmt"
)

func main() {
	var valores [10]float64

	// Leitura dos 10 valores
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº valor: ", i+1)
		fmt.Scan(&valores[i])
	}

	// Exibição em ordem inversa
	fmt.Println("\nValores na ordem inversa:")
	for i := len(valores) - 1; i >= 0; i-- {
		fmt.Printf("%.2f ", valores[i])
	}
	fmt.Println()
}
