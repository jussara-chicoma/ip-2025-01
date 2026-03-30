package main

import "fmt"

func main() {
	var preco float64
	var dia int
	var tipo string

	fmt.Scan(&preco, &dia, &tipo)

	if dia == 2 || dia == 3 || dia == 5 {
		preco *= 0.6
	}

	if tipo == "l" {
		preco *= 1.15
	}

	fmt.Println(preco)
}
