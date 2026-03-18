package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 1; i < t; i++ {
		var publico float64
		var pPopular, pGeral, pArq, pCadeira float64
		fmt.Scan(&publico, &pPopular, &pGeral, &pArq, &pCadeira)

		qPopular := float64(publico) * (pPopular / 100)
		qGeral := float64(publico) * (pGeral / 100)
		qArq := float64(publico) * (pArq / 100)
		qCadeira := float64(publico) * (pCadeira / 100)

		renda := (qPopular * 1.0) + (qGeral * 5.0) + (qArq * 10.0) + (qCadeira * 20.0)
		fmt.Printf("A renda do jogo N.%d E =%2f\n", i, renda)
	}
}
