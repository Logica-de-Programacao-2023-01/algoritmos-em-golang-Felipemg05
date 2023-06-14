package main

import "fmt"

func main() {
    var num1, num2 int

    fmt.Print("Digite o primeiro número inteiro: ")
    fmt.Scan(&num1)

    fmt.Print("Digite o segundo número inteiro: ")
    fmt.Scan(&num2)

    if num1 > 0 && num2 > 0 {
        multiplicacao := num1 * num2
        fmt.Println("Resultado da multiplicação:", multiplicacao)
    } else {
        soma := num1 + num2
        fmt.Println("Resultado da soma:", soma)
    }
}
