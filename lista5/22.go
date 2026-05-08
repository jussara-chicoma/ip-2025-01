package main

import "fmt"

func main() {
	var codigos [10]int
	var saldos [10]float64

	// leitura
	for i := 0; i < 10; i++ {
		fmt.Scan(&codigos[i], &saldos[i])
	}

	for {
		var op int
		fmt.Println("\n1 Deposito\n2 Saque\n3 Total\n4 Sair")
		fmt.Scan(&op)

		if op == 4 {
			break
		}

		var cod int
		fmt.Print("Codigo: ")
		fmt.Scan(&cod)

		pos := -1
		for i := 0; i < 10; i++ {
			if codigos[i] == cod {
				pos = i
				break
			}
		}

		if pos == -1 {
			fmt.Println("Conta nao encontrada")
			continue
		}

		if op == 1 {
			var valor float64
			fmt.Scan(&valor)
			saldos[pos] += valor
		}

		if op == 2 {
			var valor float64
			fmt.Scan(&valor)
			if saldos[pos] >= valor {
				saldos[pos] -= valor
			} else {
				fmt.Println("Saldo insuficiente")
			}
		}

		if op == 3 {
			total := 0.0
			for i := 0; i < 10; i++ {
				total += saldos[i]
			}
			fmt.Println("Total:", total)
		}
	}
}
