package main

import (
	"fmt"
)

func main() {
	var n int
	var S float64 = 0.0

	// Leitura do valor
	fmt.Scan(&n)

	// Validação
	if n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}

	// Cálculo da soma: S = 1 + 1/2 + ... + 1/n
	for k := 1; k <= n; k++ {
		S += 1.0 / float64(k)
	}

	// Saída com 6 casas decimais
	fmt.Printf("%.6f\n", S)
}
