package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Qual é sua idade:")
	fmt.Scan(&idade)
	fmt.Println("Idade em dias é:", idade*365)
}
