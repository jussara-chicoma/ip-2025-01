package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a float64
	fmt.Scan(&h, &a)

	// Área da base (hexágono regular)
	Ab := (3 * math.Pow(a, 2) * math.Sqrt(3)) / 2

	// pá calcula el Volume da pirâmide
	V := (Ab * h) / 3
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", V)
}
