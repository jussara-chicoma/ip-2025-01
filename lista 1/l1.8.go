package main

import "fmt"

func main() {
	var r, h float64
	const pi = 3.14159

	fmt.Scan(&r)
	fmt.Scan(&h)
	ac := pi * r * r
	al := 2 * pi * r * h
	at := 2*ac + al
	custo := at * 100.0
	fmt.Printf("o valor do custo e= %.2fn\"", custo)
}
