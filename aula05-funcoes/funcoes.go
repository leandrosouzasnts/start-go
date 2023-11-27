package main

import "fmt"

//Funcao que recebe dos valores do tipo int e retorna a soma deles do mesmo tipo
func somar(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(somar(1, 2))

	var f = func(text string) string { //Nesse caso func também é um tipo de dado que sua assinatura é o seu contexto (parms, return)
		return text
	}

	resulText := f("Leandro")
	fmt.Println(resulText)
}
