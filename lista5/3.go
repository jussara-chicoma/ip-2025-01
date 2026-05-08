package main

import (
	"fmt"
)

func main() {
	var vetor [10]int
	var somaPares int
	var qtdImpares int

	fmt.Println("Digite 10 números inteiros:")

	// Leitura dos valores
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	fmt.Println("\nNúmeros pares:")
	for _, v := range vetor {
		if v%2 == 0 {
			fmt.Printf("%d ", v)
			somaPares += v
		}
	}

	fmt.Println("\nSoma dos números pares:", somaPares)

	fmt.Println("\nNúmeros ímpares:")
	for _, v := range vetor {
		if v%2 != 0 {
			fmt.Printf("%d ", v)
			qtdImpares++
		}
	}

	fmt.Println("\nQuantidade de números ímpares:", qtdImpares)
}
