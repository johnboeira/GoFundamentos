package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	resultado := soma(9, 18, 32)

	fmt.Println(resultado)
}
