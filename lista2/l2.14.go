package main

import "fmt"

func main() {
	var preco float64
	var op string

	fmt.Scan(&preco)

	fmt.Scan(&op)
	if op == "s" {
		preco += 1750
	}

	fmt.Scan(&op)
	if op == "s" {
		preco += 800
	}

	fmt.Scan(&op)
	if op == "s" {
		preco += 1200
	}

	fmt.Scan(&op)
	if op == "s" {
		preco += 2000
	}

	fmt.Println(preco)
}
