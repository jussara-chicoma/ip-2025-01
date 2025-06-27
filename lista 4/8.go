package main

import (
	"fmt"
	"math"
)

func main() {
	var entrada [15]int
	var resultado [15]float64

	fmt.Println("Digite 15 números inteiros:")

	for i := 0; i < 15; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&entrada[i])

		if entrada[i] >= 0 {
			resultado[i] = math.Sqrt(float64(entrada[i]))
		} else {
			resultado[i] = -1
		}
	}

	fmt.Println("\nValores armazenados (raiz quadrada ou -1 se negativo):")
	for i := 0; i < 15; i++ {
		fmt.Printf("Posição %d: %.2f\n", i, resultado[i])
	}
}
