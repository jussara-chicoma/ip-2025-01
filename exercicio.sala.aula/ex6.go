package main

import "fmt"

func main() {
	var qtdFarinha, valorFarinha float64
	var qtdFermento, valorFermento float64
	var qtdLuz, valorLuz float64
	var imposto float64

	// Leitura dos dados
	fmt.Scan(&qtdFarinha)
	fmt.Scan(&valorFarinha)
	fmt.Scan(&qtdFermento)
	fmt.Scan(&valorFermento)
	fmt.Scan(&qtdLuz)
	fmt.Scan(&valorLuz)
	fmt.Scan(&imposto)

	// Cálculo do custo base
	custoBase := (qtdFarinha * valorFarinha) +
		(qtdFermento * valorFermento) +
		(qtdLuz * valorLuz)

	// Preço de custo (com imposto)
	precoCusto := custoBase * (1 + imposto/100)

	// Preço de venda (lucro de 30%)
	precoVenda := precoCusto * 1.30

	// Saída formatada
	fmt.Printf("PRECO DE CUSTO = %.2f\n", precoCusto)
	fmt.Printf("PRECO DE VENDA = %.2f\n", precoVenda)
}
