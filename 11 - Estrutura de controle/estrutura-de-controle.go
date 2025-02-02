package main

import "fmt"

func main() {
	numero := -2

	if numero > 15 {
		fmt.Println("Maior q 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outroNumero := numero; outroNumero <= 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Positivo")
	}
}
