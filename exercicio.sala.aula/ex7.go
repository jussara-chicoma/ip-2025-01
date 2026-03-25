package main

import "fmt"

func main() {
	var a, b, c float64
	if a >= b+c || b >= a+c || c >= a+b {
		fmt.Println("Não forma um triângulo")
		return
	}
	if a == b && b == c {
		fmt.Println("T.Equilátero")
	} else if a == b || a == c || b == c {
		fmt.Println("T.Isóscele")
	} else {
		fmt.Println("T.Escaleno")
	}
}
