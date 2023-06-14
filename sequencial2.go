package main

import "fmt"

func main() {
	var num int
	fmt.Print("")
	fmt.Print("Digite um número:")
	fmt.Scan(&num)
	fmt.Println("Dobro:", num*2)
	fmt.Println("Triplo:", num*3)
	fmt.Println("Quadruplo:", num*4)
}
