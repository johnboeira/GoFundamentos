package main

import "fmt"

func inverterValor(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	inverterValor(&numero)
	fmt.Println(numero)
}
