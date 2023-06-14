package main

import "fmt"

func main() {

	var num1, num2, num3 float64
	var peso1, peso2, peso3 float64
	peso1 = 2
	peso2 = 3
	peso3 = 5

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&num3)

	mediaPonderada := (num1*peso1 + num2*peso2 + num3*peso3) / (peso1 + peso2 + peso3)

	fmt.Println("A média ponderada é:", mediaPonderada)
}
