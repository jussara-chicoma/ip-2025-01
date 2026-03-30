package main

import "fmt"

func main() {
	var mat, horas int
	fmt.Scan(&mat, &horas)

	bruto := 3*788 + float64(horas*10)

	inss := 0.0
	ir := 0.0

	if bruto > 1500 {
		inss = bruto * 0.12
	}
	if bruto > 2000 {
		ir = bruto * 0.2
	}

	liquido := bruto - inss - ir

	fmt.Println(liquido)
}
