package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int
	fmt.Scan(&n1, &n2, &n3)

	maior := n1

	if n2 > maior {
		maior = n2
	}

	if n3 > maior {
		maior = n3
	}

	fmt.Println("O maior número é:", maior)
}
