package main

import (
	"fmt"
)

// função soma
func soma(a int, b int) int {
	return a + b
}

func main() {
	var num1, num2 int

	// leitura dos números
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	// chamada da função
	resultado := soma(num1, num2)

	// exibição do resultado
	fmt.Println("A soma é:", resultado)
}
