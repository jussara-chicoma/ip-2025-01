package main

import "fmt"

func main() {
	var n1, n2, n3, me float64
	fmt.Scan(&n1, &n2, &n3, &me)

	media := (n1 + n2*2 + n3*3 + me) / 7

	var conceito string

	if media >= 9 {
		conceito = "A"
	} else if media >= 7.5 {
		conceito = "B"
	} else if media >= 6 {
		conceito = "C"
	} else if media >= 4 {
		conceito = "D"
	} else {
		conceito = "E"
	}

	fmt.Println(media, conceito)
}
