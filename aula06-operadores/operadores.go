package main

import "fmt"

func main() {
	//ARITIMETICOS
	sub := 2 - 1
	soma := 2 + 1
	multi := 2 * 1
	div := 2 / 1
	restoDiv := 10 % 2

	fmt.Println(sub, soma, multi, div, restoDiv)

	//Atribuições
	var texto string = "String"
	texto1 := "String"

	fmt.Println(texto, texto1)

	//Relacionais
	fmt.Println(1 > 2, 2 <= 1, 2 == 2, 2 != 1)

	//Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//Operadores Unários
	numero := 10
	numero++
	numero += 20 // numero = numero + 20

	numero--
	numero -= 10 // numero = numero - 10

	numero *= 3 // numero = numero * 3

	numero /= 10 // numero = numero / 10

}
