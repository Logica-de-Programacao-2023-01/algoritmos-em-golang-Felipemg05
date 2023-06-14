package main

import "fmt"

func main() {
	var altura float64
	var peso float64
	fmt.Print("Qual é o peso?")
	fmt.Scan(&peso)
	fmt.Print("Qual é a altura?")
	fmt.Scan(&altura)
	fmt.Println("O IMC é", peso/(altura*altura))
}
