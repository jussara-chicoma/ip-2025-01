package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	binario := ""

	for n > 0 {
		resto := n % 2
		binario = fmt.Sprintf("%d", resto) + binario
		n /= 2
	}

	fmt.Println("Binário:", binario)
}
