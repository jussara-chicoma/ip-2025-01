package main

import "fmt"

func main() {
	var tipo string
	var consumo float64

	fmt.Scan(&tipo, &consumo)

	if tipo == "r" {
		fmt.Println(5 + consumo*0.05)
	} else if tipo == "c" {
		if consumo <= 80 {
			fmt.Println(500)
		} else {
			fmt.Println(500 + (consumo-80)*0.25)
		}
	} else if tipo == "i" {
		if consumo <= 100 {
			fmt.Println(800)
		} else {
			fmt.Println(800 + (consumo-100)*0.04)
		}
	}
}
