package main

import "fmt"

func closure() func() {
	texto := "Dentro da func closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "dentro da main"
	fmt.Println(texto)

	funcao := closure()
	funcao()
}
