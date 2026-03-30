package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 20 && n < 90 {
		fmt.Println("Está na faixa")
	} else {
		fmt.Println("Fora da faixa")
	}
}
