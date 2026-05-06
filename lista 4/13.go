package main

import "fmt"

type Empregado struct {
	codigo int
	meses  int
}

func main() {
	var lista []Empregado

	for {
		var c, m int
		fmt.Print("Codigo e meses (0 0 para parar): ")
		fmt.Scan(&c, &m)

		if c == 0 && m == 0 {
			break
		}

		lista = append(lista, Empregado{c, m})
	}

	//  para ordenar pelo menor tempo de trabalho
	for i := 0; i < len(lista); i++ {
		for j := i + 1; j < len(lista); j++ {
			if lista[j].meses < lista[i].meses {
				lista[i], lista[j] = lista[j], lista[i]
			}
		}
	}

	fmt.Println("\nMais recentes:")
	for i := 0; i < 3 && i < len(lista); i++ {
		fmt.Println(lista[i])
	}
}
