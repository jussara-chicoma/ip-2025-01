package main

import "fmt"

func main() {
	var vet [100]int

	// Ler 100 valores numéricos
	fmt.Println("Digite 100 números:")
	for i := 0; i < 100; i++ {
		fmt.Scan(&vet[i])
	}

	// Calcular o somatório conforme a fórmula fornecida
	var soma int
	for i := 0; i < 50; i++ {
		soma += (vet[i] - vet[99-i]) * (vet[i] - vet[99-i]) * (vet[i] - vet[99-i])
	}

	fmt.Printf("O valor do somatório é: %d\n", soma)
}
