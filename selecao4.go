package main

import "fmt"

func main() {
    var altura, peso float64
    var sexo string

    fmt.Print("Digite a altura em metros: ")
    fmt.Scan(&altura)

    fmt.Print("Digite o peso em quilos: ")
    fmt.Scan(&peso)

    fmt.Print("Digite o sexo (M para masculino, F para feminino): ")
    fmt.Scan(&sexo)

    var k float64

    if sexo == "M" {
        k = 4
    } else if sexo == "F" {
        k = 2
    } else {
        fmt.Println("Sexo inválido.")
        return
    }

    pesoIdeal := altura - k

    if peso < pesoIdeal-2 {
        fmt.Println("Abaixo do peso ideal.")
    } else if peso <= pesoIdeal+2 {
        fmt.Println("Dentro do peso ideal.")
    } else {
        fmt.Println("Acima do peso ideal.")
    }
}
