package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println()

	var u usuario
	u.nome = "dave"
	u.idade = 21
	fmt.Println(u)

	endereco := endereco{"Rua", 0}

	//inferencia
	u2 := usuario{"Davi", 21, endereco}
	fmt.Println(u2)

	u3 := usuario{nome: "Dave"}
	fmt.Println(u3)
}
