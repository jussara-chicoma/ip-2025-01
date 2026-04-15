package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4, 5}

	n := len(v)

	for i := 0; i < n/2; i++ {
		v[i], v[n-1-i] = v[n-1-i], v[i]
	}

	fmt.Println(v)
}
