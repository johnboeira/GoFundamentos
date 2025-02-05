package main

import "fmt"

func verificaAprovacao(nota1, nota2 float32) bool {
	defer fmt.Println("Media calculada")

	media := (nota1 + nota2) / 2

	if media > 7 {
		return true
	}

	return false

}

func main() {
	aprovado := verificaAprovacao(5, 5)

	if aprovado {
		fmt.Println("Aprovado")
	}

	fmt.Println("Reprovado")
}
