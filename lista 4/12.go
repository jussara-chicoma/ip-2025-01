package main

import "fmt"

func main() {
	var notas [15]int
	var freqAbs [11]int

	// Ler as 15 notas dos alunos
	fmt.Println("Digite as 15 notas dos alunos (de 0 a 10):")
	for i := 0; i < 15; i++ {
		fmt.Scan(&notas[i])
		freqAbs[notas[i]]++
	}

	// Calcular a frequência relativa
	var freqRel [11]float64
	for i := 0; i <= 10; i++ {
		freqRel[i] = float64(freqAbs[i]) / 15
	}

	// Imprimir a tabela de frequências
	fmt.Println("\nNota | Frequência Absoluta | Frequência Relativa")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%2d   | %d                    | %.2f\n", i, freqAbs[i], freqRel[i])
	}
}
