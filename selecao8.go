package main

import "fmt"

func main() {
    var numero int

    fmt.Print("Digite um número inteiro entre 1 e 7: ")
    fmt.Scan(&numero)

    var diaSemana string

    switch numero {
    case 1:
        diaSemana = "Domingo"
    case 2:
        diaSemana = "Segunda-feira"
    case 3:
        diaSemana = "Terça-feira"
    case 4:
        diaSemana = "Quarta-feira"
    case 5:
        diaSemana = "Quinta-feira"
    case 6:
        diaSemana = "Sexta-feira"
    case 7:
        diaSemana = "Sábado"
    default:
        diaSemana = "Número inválido"
    }

    fmt.Println("Dia da semana correspondente:", diaSemana)
}
