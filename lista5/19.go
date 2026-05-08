package main

import "fmt"

func main() {
	num := [10]int{5, 12, 4, 7, 10, 3, 2, 6, 23, 16}
	div := [5]int{3, 11, 5, 8, 2}

	for i := 0; i < 10; i++ {
		fmt.Println("Numero", num[i])
		for j := 0; j < 5; j++ {
			if num[i]%div[j] == 0 {
				fmt.Println("Divisivel por", div[j], "na posicao", j)
			}
		}
	}
}