package main

import "fmt"

func main() {
	var v1, v2 [10]int
	var v3 [20]int

	for i := 0; i < 10; i++ {
		fmt.Scan(&v1[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Scan(&v2[i])
	}

	k := 0
	for i := 0; i < 10; i++ {
		v3[k] = v1[i]
		k++
		v3[k] = v2[i]
		k++
	}

	fmt.Println(v3)
}