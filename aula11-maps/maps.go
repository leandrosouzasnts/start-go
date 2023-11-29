package main

import "fmt"

func main() {

	usuario1 := map[string]string{
		"nome":      "Leandro",
		"sobrenome": "Souza",
	}
	fmt.Println(usuario1["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Leandro",
			"segundo":  "Souza",
		},
		"curso": {
			"nome": "Ads",
		},
	}
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{ //Adicionando
		"nome": "Capricornio",
	}
	fmt.Println(usuario2)

	delete(usuario2, "signo") //Remover
	fmt.Println(usuario2)

}
