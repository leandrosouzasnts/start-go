package main

import "fmt"

func main() {

	var var1 int = 10
	var var2 int = var1

	var1++

	fmt.Println(var1, var2)

	var ponteiro *int //o tipo do ponteiro deve ser do mesmo tipo que a referencia (ou seja, valor guardado nesse endereço)
	ponteiro = &var1

	var1--

	fmt.Println(var1, *ponteiro) //desreferenciação (*ponteiro está buscando o valor no mesmo local que var1)
}
