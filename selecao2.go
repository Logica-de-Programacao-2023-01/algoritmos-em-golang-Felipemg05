package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64
	fmt.Print("Qual é o valor x ?")
	fmt.Scan(&x)
	fmt.Println("X:", x)
	fmt.Print("Qual é o valor de y?")
	fmt.Scan(&y)
	fmt.Println("Y:", y)
	fmt.Print("Qual é o valor de Z?")
	fmt.Scan(&z)
	fmt.Println("Z:", z)

	if x < y && x < z {
		fmt.Println("x é menor que y e z ")
	} else if y < x && y < z {
		fmt.Println("y é menor que x e z")
	} else if z < x && z < y {
		fmt.Println("z é menor que x e y")
	}
}
