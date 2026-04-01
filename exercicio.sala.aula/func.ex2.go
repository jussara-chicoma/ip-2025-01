package main

import (
	"fmt"
)

// função que retorna o maior entre três números
func maior(a int, b int, c int) int {
	maior := a

	if b > maior {
		maior = b
	}

	if c > maior {
		maior = c
	}

	return maior
}

func main() {
	var num1, num2, num3 int

	// leitura dos valores
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)

	// chamada da função
	resultado := maior(num1, num2, num3)

	// exibição do resultado
	fmt.Println("O maior número é:", resultado)
}
