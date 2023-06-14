package main

import "fmt"

func main() {
    var idade int

    fmt.Print("Digite a idade da pessoa: ")
    fmt.Scan(&idade)

    var classificacao string

    if idade <= 9 {
        classificacao = "Mirim"
    } else if idade <= 13 {
        classificacao = "Infantil"
    } else if idade <= 17 {
        classificacao = "Juvenil"
    } else {
        classificacao = "Adulto"
    }

    fmt.Println("Classificação:", classificacao)
}
