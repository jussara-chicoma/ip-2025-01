package main

import (
	"fmt"
)

func main() {
	var valores [5]int
	soma := 0

	// Leitura dos 5 valores
	for i := 0; i < 5; i++ {
		fmt.Printf("Digite o %dº valor: ", i+1)
		fmt.Scan(&valores[i])
		soma += valores[i]
	}

	// Exibe o vetor (opcional)
	fmt.Println("\nValores digitados:", valores)

	// Exibe a soma
	fmt.Println("Soma dos valores:", soma)
}
