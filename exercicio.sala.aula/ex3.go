package main

import (
	"fmt"
)

func main() {
	var valor float64
	var percentual float64
	fmt.Print("digite o valor da mercadoria:")
	fmt.Scan(&valor)
	fmt.Print("digite o percentual de deconto:")
	fmt.Scan(&percentual)

	desconto := valor * (percentual / 100)
	novoValor := valor - desconto
	fmt.Println("o valor do desconto:", desconto)
	fmt.Println("o novo valor da mercadoria:", novoValor)

}
