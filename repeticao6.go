package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número :", num)
	fmt.Scan(&num)

	fmt.Println("Tabuada do número", num, ":")

	for i := 1; i <= 10; i++ {
		resultado := num * i
		fmt.Println(num, "x", i, "=", resultado)
	}
}
