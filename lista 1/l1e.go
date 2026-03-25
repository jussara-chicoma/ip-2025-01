package main

import (
	"fmt"
)

func main() {
	var salarioMinimo float64
	var consumoKW float64

	// Entrada
	fmt.Scan(&salarioMinimo)
	fmt.Scan(&consumoKW)

	// Cálculos
	custoPorKW := (salarioMinimo * 0.7) / 100
	custoConsumo := consumoKW * custoPorKW
	custoDesconto := custoConsumo * 0.9

	// Saída formatada
	fmt.Printf("Custo por kW: R$ %.2f\n", custoPorKW)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}
