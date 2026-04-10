package main

import "fmt"

func main() {
	var base, expoente int
	resultado := 1
	fmt.Println("digite a base, ")
	fmt.Scan(&base)
	fmt.Print("digite o expoente ")
	fmt.Scan(&expoente)
	for i := 0; i < expoente; i++ {
		resultado *= base
	}
	fmt.Printf("resultado :%d\n", resultado)

}
