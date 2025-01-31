package main

import "fmt"

func main() {
	var variavel1 int = 100
	var ponteiro *int

	ponteiro = &variavel1

	fmt.Println(variavel1, ponteiro)
	fmt.Println(variavel1, *ponteiro)

	variavel1 = 150

	fmt.Println(variavel1, *ponteiro)
}
