package main

import (
	"fmt"
)

func main() {
	var numero float64
	var soma float64

	for i := 1; i <= 10; i++ {
		fmt.Printf("Digite o %dº número: ", i)
		fmt.Scan(&numero)

		soma += numero
	}

	media := soma / 10

	fmt.Printf("A média aritmética é: %.2f\n", media)
}
