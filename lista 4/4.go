package main

import (
	"fmt"
)

func main() {
	var A [10]int
	frequencia := make(map[int]int)

	// Leitura do vetor
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&A[i])
		frequencia[A[i]]++
	}

	// Verificando e mostrando os elementos repetidos
	fmt.Println("\nElementos repetidos e suas frequências:")
	encontrouRepetido := false
	for valor, freq := range frequencia {
		if freq > 1 {
			fmt.Printf("Valor %d aparece %d vezes\n", valor, freq)
			encontrouRepetido = true
		}
	}

	if !encontrouRepetido {
		fmt.Println("Nenhum elemento repetido encontrado.")
	}
}
