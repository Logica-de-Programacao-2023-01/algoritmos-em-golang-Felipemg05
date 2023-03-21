package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Qual o valor do número?")
	fmt.Scan(&numero)
	fmt.Println("Número:", numero)

	if numero%2 == 0 {
		fmt.Println("O número é par")
	} else if numero%2 == 1 {
		fmt.Println("O número é ímpar")
	}

}
