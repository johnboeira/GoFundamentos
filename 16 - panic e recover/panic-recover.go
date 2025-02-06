package main

import "fmt"

func recuperacao() {
	if r := recover(); r != nil {
		fmt.Println("Exec recuperada")
	}
}

func verificaAprovacao(n1, n2 float32) bool {
	defer recuperacao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("média é 6!")
}

func main() {
	verificaAprovacao(6, 6)
	fmt.Println("Pos exec")
}
