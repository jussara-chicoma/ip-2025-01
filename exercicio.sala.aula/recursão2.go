package main

import "fmt"

func main() {
	var n int
	fmt.Print("Tamanho: ")
	fmt.Scan(&n)

	valores := make([]float64, n)
	soma := 0.0

	for i := 0; i < n; i++ {
		fmt.Scan(&valores[i])
		soma += valores[i]
	}

	fmt.Println("Soma:", soma)
}
