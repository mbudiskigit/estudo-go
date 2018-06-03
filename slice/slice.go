package main

import "fmt"

func main() {

	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", original)
	novo := original[1:3]
	fmt.Println("Novo:", novo)
	original[2] = 13
	fmt.Println("Original pós modificação:", original)
	fmt.Println("Novo pós modificação:", novo)
}
