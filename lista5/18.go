package main

import "fmt"

func main() {
	var v [10]int

	for i := 0; i < 10; i++ {
		for {
			fmt.Scan(&v[i])
			if i == 0 || v[i] >= v[i-1] {
				break
			}
			fmt.Println("Digite um valor maior ou igual ao anterior")
		}
	}

	fmt.Println(v)
}