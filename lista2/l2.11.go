package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	if x == 2 {
		fmt.Println("Indefinido")
	} else {
		fmt.Println(8 / (2 - x))
	}
}
