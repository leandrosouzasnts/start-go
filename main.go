package main

import (
	"fmt"
	"modulo/aula01-pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do package main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("leandro@teste.com.br")

	fmt.Println(erro)

}
