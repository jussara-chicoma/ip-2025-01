package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int
	fmt.Scan(&n1, &n2, &n3)

	if n1 > n2 {
		n1, n2 = n2, n1
	}
	if n1 > n3 {
		n1, n3 = n3, n1
	}
	if n2 > n3 {
		n2, n3 = n3, n2
	}

	fmt.Println(n1, n2, n3)
}
