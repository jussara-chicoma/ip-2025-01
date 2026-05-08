package main

import "fmt"

func main() {
	var v [50]int
	freq := make(map[int]int)

	for i := 0; i < 50; i++ {
		fmt.Scan(&v[i])
		freq[v[i]]++
	}

	moda := v[0]
	maior := 0

	for k, val := range freq {
		if val > maior {
			maior = val
			moda = k
		}
	}

	fmt.Println("Moda:", moda)
}