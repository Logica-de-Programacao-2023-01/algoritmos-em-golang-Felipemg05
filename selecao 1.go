package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Qual o valor de x?")
	fmt.Scan(&x)
	fmt.Println("X:")
	fmt.Print("Qual o valor de y?")
	fmt.Scan(&y)
	fmt.Println("Y:")

	if x > y {

		fmt.Println("x é maior que y")

	} else {

		fmt.Println("y é maior que x")

	}
}
