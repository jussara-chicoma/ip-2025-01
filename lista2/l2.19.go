package main

import (
	"fmt"
	"math"
)

func main() {
	var op int
	var r, h float64

	fmt.Scan(&op)
	fmt.Scan(&r)

	if op != 3 {
		fmt.Scan(&h)
	}

	if op == 1 {
		v := math.Pi * r * r * h / 3
		a := math.Pi * r * math.Sqrt(r*r+h*h)
		fmt.Println(v, a)
	} else if op == 2 {
		v := math.Pi * r * r * h
		a := 2 * math.Pi * r * h
		fmt.Println(v, a)
	} else if op == 3 {
		v := 4.0 / 3.0 * math.Pi * r * r * r
		a := 4 * math.Pi * r * r
		fmt.Println(v, a)
	}
}
