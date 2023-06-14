package main

import "fmt"

func main() {
    var preco float64

    fmt.Print("Digite o preço do produto: ")
    fmt.Scan(&preco)

    desconto := preco * 0.1
    precoComDesconto := preco - desconto

    fmt.Println("O preço com desconto é:", precoComDesconto)
}
