package main

import "fmt"

func main() {
	//Arit
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//Relacional
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//Logicos
	fmt.Println("------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)

	//Unitarios
	var numero float32 = 10
	numero++
	fmt.Println(numero)
	numero += 10
	fmt.Println(numero)
	numero /= 2
	fmt.Println(numero)

}
