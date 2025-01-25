package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int8 = 100
	fmt.Println(numero)

	//int32
	var numero2 rune = 100
	fmt.Println(numero2)

	var numeroint8 byte = 123
	fmt.Println(numeroint8)

	//FLOAT
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	//CHAR int de ASCII
	char := 'B'
	fmt.Println(char)

	//Bool
	var boolean1 bool = false
	fmt.Println(boolean1)

	//Erro
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro novo")
	fmt.Println(erro2)
}
