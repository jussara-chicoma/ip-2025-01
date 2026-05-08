package main

import (
	"fmt"
)

func main() {
	var alturas [10]float64
	var soma float64

	// Leitura das alturas
	fmt.Println("Digite a altura de 10 atletas (em metros, ex: 1.75):")
	for i := 0; i < 10; i++ {
		fmt.Printf("Altura do atleta %d: ", i+1)
		fmt.Scan(&alturas[i])
		soma += alturas[i]
	}

	// Cálculo da média
	media := soma / 10

	fmt.Printf("\nAltura média: %.2f metros\n", media)
	fmt.Println("Atletas com altura acima da média:")
	for _, altura := range alturas {
		if altura > media {
			fmt.Printf("%.2f metros\n", altura)
		}
	}
}
