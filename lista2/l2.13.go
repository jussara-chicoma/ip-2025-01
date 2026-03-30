package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n >= 100 && n <= 999 {
		fmt.Println((n / 10) % 10)
	} else {
		fmt.Println("Inválido")
	}
}
