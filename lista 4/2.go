package main

import (
	"fmt"
)

func main() {
	var vetor1 [10]int
	var vetor2 [5]int
	var resultPares []int
	var resultImpares []int
	var somaVetor2 int

	// Leitura do primeiro vetor
	fmt.Println("Digite 10 números inteiros para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	// m´ta face Leitura do segundo vetor
	fmt.Println("\nDigite 5 números inteiros para o segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor2[i])
		somaVetor2 += vetor2[i]
	}

	// aqui m´ta Separa os pares e ímpares mode m´ta face  com soma
	for _, num := range vetor1 {
		if num%2 == 0 {
			resultPares = append(resultPares, num+somaVetor2)
		} else {
			resultImpares = append(resultImpares, num+somaVetor2)
		}
	}

	// Exibindo resultados
	fmt.Println("\nPrimeiro vetor resultante (pares + soma do segundo vetor):", resultPares)
	fmt.Println("Segundo vetor resultante (ímpares + soma do segundo vetor):", resultImpares)
}
