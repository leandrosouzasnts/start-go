package main

import "fmt"

//"Herança"

type Pessoa struct {
	nome  string
	idade uint8
}

type Estudante struct {
	Pessoa
	curso string
}

func main() {

	pessoa := Pessoa{"Leandro", 26}
	estudante := Estudante{pessoa, "ADS"}

	fmt.Println(estudante.nome)

}
