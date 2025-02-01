package main

import "fmt"

func main() {
	usuario := map[string]string{
		"Nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"Nome": {
			"Primeiro ": "João",
			"Ultimo":    "Pedro",
		},
		"Curso": {
			"Nome":   "Eng",
			"Campus": "Capm 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "Nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}

	fmt.Println(usuario2)

}
