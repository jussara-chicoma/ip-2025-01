package main

import (
	"fmt"
)

func main() {
	var vetor [10]int
	var menor, posicao int

	fmt.Println("Digite 10 números inteiros:")

	
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
	}


	menor = vetor[0]
	posicao = 0


	for i := 1; i < 10; i++ {
		if vetor[i] < menor {
			menor = vetor[i]
			posicao = i
		}
	}

	// Resultado
	fmt.Printf("\nO menor elemento do vetor é %d e sua posição é %d\n", menor, posicao)
}
