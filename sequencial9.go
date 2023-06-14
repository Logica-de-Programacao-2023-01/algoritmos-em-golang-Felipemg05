package main

import "fmt"

func main() {
    var pesoEmQuilos float64

    fmt.Print("Digite o peso em quilos: ")
    fmt.Scan(&pesoEmQuilos)

    pesoEmLibras := pesoEmQuilos * 2.20462

    fmt.Println("O peso em libras é:", pesoEmLibras)
}
