package main

import "fmt"

func main() {
	var salarioCarlos float64

	fmt.Print("Digite o salário de Carlos: ")
	fmt.Scan(&salarioCarlos)

	salarioJoao := salarioCarlos / 3

	carlos := salarioCarlos
	joao := salarioJoao

	meses := 0

	for joao < carlos {
		carlos *= 1.02
		joao *= 1.05
		meses++
	}

	fmt.Println("Meses necessários:", meses)
}
