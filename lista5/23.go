package main

import (
	"fmt"
)

func mostrarDisponiveis(poltronas []int, tipo string) {
	fmt.Printf("Poltronas disponíveis (%s):\n", tipo)
	for i := 0; i < len(poltronas); i++ {
		if poltronas[i] == 0 {
			fmt.Printf(" %d", i+1)
		}
	}
	fmt.Println()
}

func todasOcupadas(poltronas []int) bool {
	for _, v := range poltronas {
		if v == 0 {
			return false
		}
	}
	return true
}

func main() {
	var janela [24]int
	var corredor [24]int
	var opcao, escolha int

	for {
		fmt.Println("\n=== Venda de Passagens ===")
		if todasOcupadas(janela[:]) && todasOcupadas(corredor[:]) {
			fmt.Println("Ônibus completamente lotado.")
			break
		}

		fmt.Println("Escolha o tipo de poltrona:")
		if !todasOcupadas(janela[:]) {
			fmt.Println("1 - Janela")
		}
		if !todasOcupadas(corredor[:]) {
			fmt.Println("2 - Corredor")
		}
		fmt.Print("Opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			if todasOcupadas(janela[:]) {
				fmt.Println("Não há mais poltronas na janela.")
				continue
			}
			mostrarDisponiveis(janela[:], "janela")
			fmt.Print("Escolha o número da poltrona (1 a 24): ")
			fmt.Scan(&escolha)
			if escolha < 1 || escolha > 24 {
				fmt.Println("Poltrona inválida.")
			} else if janela[escolha-1] == 1 {
				fmt.Println("Poltrona já ocupada.")
			} else {
				janela[escolha-1] = 1
				fmt.Println("Reserva efetuada com sucesso!")
			}
		case 2:
			if todasOcupadas(corredor[:]) {
				fmt.Println("Não há mais poltronas no corredor.")
				continue
			}
			mostrarDisponiveis(corredor[:], "corredor")
			fmt.Print("Escolha o número da poltrona (1 a 24): ")
			fmt.Scan(&escolha)
			if escolha < 1 || escolha > 24 {
				fmt.Println("Poltrona inválida.")
			} else if corredor[escolha-1] == 1 {
				fmt.Println("Poltrona já ocupada.")
			} else {
				corredor[escolha-1] = 1
				fmt.Println("Reserva efetuada com sucesso!")
			}
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
