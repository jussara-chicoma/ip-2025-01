package main

import (
	"fmt"
)

func main() {
	var salario float64
	var reajustado float64

	// Leitura do salário
	fmt.Scan(&salario)

	// Aplicação do reajuste
	if salario <= 300.00 {
		reajustado = salario * 1.50
	} else {
		reajustado = salario * 1.30
	}

	// Saída formatada
	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", reajustado)
}
