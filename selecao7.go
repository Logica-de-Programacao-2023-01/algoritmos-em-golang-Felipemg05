package main

import "fmt"

func main() {
    var salario float64

    fmt.Print("Digite o salário do funcionário: ")
    fmt.Scan(&salario)

    var aumento float64

    if salario < 1000.0 {
        aumento = salario * 0.10
    } else {
        aumento = salario * 0.05
    }

    novoSalario := salario + aumento

    fmt.Println("Novo salário do funcionário:", novoSalario)
}
