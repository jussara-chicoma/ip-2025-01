package main

import "fmt"

func main() {
	var jogadas [20]int
	var freq [7]int

	for i := 0; i < 20; i++ {
		fmt.Scan(&jogadas[i])
		freq[jogadas[i]]++
	}

	for i := 1; i <= 6; i++ {
		fmt.Println(i, "->", freq[i])
	}
}
