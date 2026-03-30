package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := make([]int, 3)
	fmt.Scan(&nums[0], &nums[1], &nums[2])

	sort.Ints(nums)

	fmt.Println("MENOR:", nums[0])
	fmt.Println("INTER:", nums[1])
	fmt.Println("MAIOR:", nums[2])
}
