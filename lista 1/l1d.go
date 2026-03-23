package main

import (
	"fmt"
)

func main() {
	var salarioMinimo float64
	var consumoKW float64

	//  nha Entrada
	fmt.Scan(&salarioMinimo)
	fmt.Scan(&consumoKW)

	//  nhos Cálculos
	custoPorKW := (salarioMinimo * 0.7) / 100
	custoConsumo := consumoKW * custoPorKW
	custoDesconto := custoConsumo * 0.9

	//  para Saída formatada sair formatada printf,2f indicar que é duas casas decimais
	fmt.Printf("Custo por kW: R$ %.2f\n", custoPorKW)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}
