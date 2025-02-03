package main

import "fmt"

func verificaDiaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Doming"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Inválido"
	}

}

func main() {
	dia := verificaDiaDaSemana(3)
	fmt.Println(dia)
}
