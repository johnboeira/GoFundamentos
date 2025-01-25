package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func Escrever() {
	fmt.Println("Escrevendo do arquivo aux")
	escrever2()
	erro := checkmail.ValidateFormat("teste  naoehvalido")

	fmt.Println(erro)
}
