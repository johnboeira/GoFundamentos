package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	resultado1, resultado2 := calculosMatematicos(10, 20)
	fmt.Println(resultado1, resultado2)
}
