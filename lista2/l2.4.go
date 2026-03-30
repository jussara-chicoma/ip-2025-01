package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	if x >= 0 {
		fmt.Println(math.Sqrt(x))
	} else {
		fmt.Println(x * x)
	}
}
