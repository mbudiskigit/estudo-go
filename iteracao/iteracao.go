package main

import "fmt"

func main() {

	a, b := 0, 10
	for a < b {
		a++
	}
	fmt.Println(a)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var x int
	for x = 0; x < 10; x++ {
		fmt.Println(x)
	}
	x++
	fmt.Println(x)

	numeros := []int{1, 2, 3, 4, 5}
	for i := range numeros {
		numeros[i] *= 2
	}

	var i int
loop:
	for i = 0; i < 10; i++ {
		fmt.Printf("for i = %d\n", i)
		switch i {
		case 2, 3:
			fmt.Printf("Quebrando switch, i == %d.\n", i)
			break
		case 5:
			fmt.Println("Quebrando loop, i == 5.")
			break loop
		}
	}
	fmt.Println("Fim.")
}
