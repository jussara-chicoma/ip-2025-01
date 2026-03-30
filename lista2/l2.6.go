package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b != 0 && a%b == 0 {
		fmt.Println("Divisível")
	} else {
		fmt.Println("Não divisível")
	}
}
