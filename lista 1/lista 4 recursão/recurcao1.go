package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x, &n)

	resultado := 1

	for i := 0; i < n; i++ {
		resultado *= x
	}

	fmt.Println("Resultado:", resultado)
}
