package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var j float64
		fmt.Scan(&j)

		c := 5 * (j - 32) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", j, c)
	}
}
