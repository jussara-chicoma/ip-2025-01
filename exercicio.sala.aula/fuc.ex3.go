package main

import "fmt"

// função MEDIA
func MEDIA(x float64, y float64, z float64) float64 {
	return (x + y + z) / 3
}

func main() {
	var x, y, z float64

	// leitura dos valores
	fmt.Print("Digite o primeiro valor: ")
	fmt.Scan(&x)

	fmt.Print("Digite o segundo valor: ")
	fmt.Scan(&y)

	fmt.Print("Digite o terceiro valor: ")
	fmt.Scan(&z)

	// chamada da função
	resultado := MEDIA(x, y, z)

	// exibição do resultado
	fmt.Println("A média é:", resultado)
}
