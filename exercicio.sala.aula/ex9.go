package main

import (
	"fmt"
)

func main() {
	var notas [3]float64
	var soma float64

	//  l.das 3 notas
	for i := 0; i < 3; i++ {
		fmt.Printf("Digite a %d nota: ", i+1)
		fmt.Scan(&notas[i])
		soma += notas[i]
	}

	// C.média
	media := soma / float64(len(notas))

	fmt.Printf("\nMédia geral: %.2f\n", media)

	// v. notas acima da média
	fmt.Print("Notas acima da média: ")
	for _, nota := range notas {
		if nota > media {
			fmt.Printf("%.2f ", nota)
		}
	}
	fmt.Println()
}
