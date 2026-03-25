package main

import (
	"fmt"
)

func main() {
	//n para representar o numero de alunos
	var n int

	var soma, nota float64
	//loop para ler todas as notas
	for i := 0; i < n; i++ {
		fmt.Scan(&nota)
		//somar as notas e acumular
		soma += nota
	}
	// calcular a media
	media := soma / float64(n)

	fmt.Printf("MEDIA = %.2f\n", media)
}
