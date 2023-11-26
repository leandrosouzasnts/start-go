package main

import "fmt"

func main() {
	//Maneiras de declarar

	//1 Forma explicita
	var var1 string
	var1 = "divani"

	//2 Forma implicita
	var var2 string = "leandro" //Comum

	//3 Forma implicita //Comum
	var3 := "carlos"

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)

	//4
	var (
		var4 string = "var4"
		var5 string = "var5"
	)

	fmt.Println(var4, var5)

	var6, var7 := "var6", "var7"

	fmt.Println(var6, var7)

	//Constante segue o mesmo modelo de todas as declarações
	const varConst string = "Const"
	fmt.Println(varConst)

	//Go permite de maneira simples trocarmos o valor de uma variável sem a necessidade de um AUX
	var6, var7 = var7, var6
	fmt.Println(var6, var7)
}
