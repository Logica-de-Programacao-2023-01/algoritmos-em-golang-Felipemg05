package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Qual é o sálario? :")
	fmt.Scanln(&salario)

	aumento := salario * 0.15
	novoSalario := salario + aumento

	fmt.Println("O novo sálario com 15% é:", novoSalario)
}
