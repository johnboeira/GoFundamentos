package main

import "fmt"

func main() {
	var variavelTexto string = "Var 1"
	//Implicito, inferencia
	variavel2 := "Var 2"
	fmt.Println(variavelTexto)
	fmt.Println(variavel2)

	var (
		variavel3 string = "a 3"
		variavel4 string = "a 4"
	)

	const cont1 string = "const 1"
	fmt.Println(cont1)

	var5, var6 := "Var 5", "Var 6"

	fmt.Println(variavel3, variavel4, var5, var6)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}
