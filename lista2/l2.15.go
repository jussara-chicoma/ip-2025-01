package main

import "fmt"

func main() {
	var d, m, a int
	fmt.Scan(&d, &m, &a)

	meses := []string{"", "janeiro", "fevereiro", "março", "abril", "maio",
		"junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}

	fmt.Printf("%d de %s de %d\n", d, meses[m], a)
}
