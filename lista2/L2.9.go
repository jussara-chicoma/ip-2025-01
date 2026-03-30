package main

import "fmt"

func main() {
	var valor float64
	fmt.Scan(&valor)

	if valor < 0 {
		fmt.Println("Valor inválido")
	} else if valor < 10 {
		fmt.Println(valor * 1.7)
	} else if valor < 30 {
		fmt.Println(valor * 1.5)
	} else if valor < 50 {
		fmt.Println(valor * 1.4)
	} else {
		fmt.Println(valor * 1.3)
	}
}
