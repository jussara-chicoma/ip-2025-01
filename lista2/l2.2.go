package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println("Positivo")
	} else if n < 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Nulo")
	}
}
