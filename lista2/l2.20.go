package main

import "fmt"

func main() {
	var preco float64
	var cod int

	fmt.Scan(&preco, &cod)

	if cod == 1 {
		preco *= 0.9
	} else if cod == 2 {
		preco *= 0.95
	} else if cod == 4 {
		preco *= 1.1
	}

	fmt.Println(preco)
}
