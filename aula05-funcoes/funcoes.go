package main

import "fmt"

//Funcao que recebe dos valores do tipo int e retorna a soma deles do mesmo tipo
func somar(n1 int, n2 int) int {
	return n1 + n2
}

//Funcao em go podem ter mais que um retorno
func calculosMatematicos(n1, n2 int) (int, int) {

	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	fmt.Println(somar(1, 2))

	var f = func(text string) string { //Nesse caso func também é um tipo de dado que sua assinatura é o seu contexto (parms, return)
		return text
	}

	resulText := f("Leandro")
	fmt.Println(resulText)

	resultadoSoma, _ := calculosMatematicos(1, 2) //Quando não queremos obter algum dos valores usamos o ' _ '
	resultadoSoma1, resultadoSub1 := calculosMatematicos(1, 2)
	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSoma1, resultadoSub1)

}
