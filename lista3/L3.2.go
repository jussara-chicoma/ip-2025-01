package main

import (
	"fmt"
)

func main() {
	soma := 0
	quantidade := 0

	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			soma += i
			quantidade++
		}
	}

	media := float64(soma) / float64(quantidade)

	fmt.Println("Soma dos valores pares:", soma)
	fmt.Println("Média dos valores pares:", media)
}
