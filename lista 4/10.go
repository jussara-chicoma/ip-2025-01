package main

import "fmt"

func main() {
	// Criar um vetor para armazenar os 50 primeiros termos da série de Fibonacci
	var fib [50]int
	fib[0] = 1
	fib[1] = 1

	for i := 2; i < 50; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println("Os primeiros 50 termos da série de Fibonacci são:")
	for i := 0; i < 50; i++ {
		fmt.Printf("%d ", fib[i])
	}
	fmt.Println()
}
