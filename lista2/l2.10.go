package main

import "fmt"

func main() {
	var destino, tipo int
	fmt.Scan(&destino, &tipo)

	var preco float64

	switch destino {
	case 1:
		if tipo == 1 {
			preco = 900
		} else {
			preco = 500
		}
	case 2:
		if tipo == 1 {
			preco = 650
		} else {
			preco = 350
		}
	case 3:
		if tipo == 1 {
			preco = 600
		} else {
			preco = 350
		}
	case 4:
		if tipo == 1 {
			preco = 550
		} else {
			preco = 300
		}
	default:
		fmt.Println("Inválido")
		return
	}

	fmt.Println(preco)
}
