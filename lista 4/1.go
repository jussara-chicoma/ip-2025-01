package main

import (
    "fmt"
)

func main() {
    var vetor [10]int
    var encontrou bool

    // Leitura dos 10 números
    fmt.Println("Digite 10 números inteiros:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Digite o %dº número: ", i+1)
        fmt.Scan(&vetor[i])
    }

    // Verificação dos números maiores que 50
    fmt.Println("\nNúmeros superiores a 50 e suas respectivas posições:")
    for i, v := range vetor {
        if v > 50 {
            fmt.Printf("Valor %d na posição %d\n", v, i)
            encontrou = true
        }
    }

    //  se nenhum número maior que 50 for encontrado
    if !encontrou {
        fmt.Println("Nenhum número superior a 50 foi encontrado.")
    }
}