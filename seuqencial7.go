ackage main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número:")
	fmt.Scan(&num)

	antecessor := num - 1
	sucessor := num + 1
	fmt.Println("Antecessor:", antecessor)
	fmt.Println("Sucessor:", sucessor)
}
