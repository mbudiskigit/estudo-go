package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for indice, valor := range entrada {

		numero, err := strconv.Atoi(valor)
		if err != nil {
			fmt.Printf("%s não é um número válido!\n", valor)
			os.Exit(1)
		}

		numeros[indice] = numero
	}

	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {

	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))

	copy(n, numeros)

	indicePivo := len(n) / 2
	pivo := n[indicePivo]

	n = append(n[:indicePivo], n[indicePivo+1:]...)

	menores, maiores := particionar(n, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {

	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}
