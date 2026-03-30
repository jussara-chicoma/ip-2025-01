package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	soma := a + b

	if soma > 20 {
		soma += 8
	} else {
		soma -= 5
	}

	fmt.Println(soma)
}
