package main

import (
	"fmt"
)

func main() {
	var vetor [100]int
	numero := 1 // primeiro número ímpar

	// Armazenando os 100 primeiros ímpares
	for i := 0; i < 100; i++ {
		vetor[i] = numero
		numero += 2 // próximo ímpar
	}

	// Imprimindo o vetor
	fmt.Println("Os 100 primeiros números ímpares são:")
	for i := 0; i < 100; i++ {
		fmt.Print(vetor[i], " ")
	}
	fmt.Println()
}
