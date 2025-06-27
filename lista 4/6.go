package main

import (
	"fmt"
)

func main() {
	var vetor [100]int

	// Armazenar números de 100 a 1
	for i := 0; i < 100; i++ {
		vetor[i] = 100 - i
	}

	// Imprimir o vetor
	fmt.Println("Números de 100 a 1 em ordem decrescente:")
	for i := 0; i < 100; i++ {
		fmt.Print(vetor[i], " ")
	}
	fmt.Println()
}
