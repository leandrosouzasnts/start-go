package main

import (
	"errors"
	"fmt"
)

func main() {

	//Valores padrão: int, float = 0/ string = "", bool = false, error = nil

	/* Error */
	var erro error = errors.New("Error")
	fmt.Println(erro)

	/* Inteiro */
	//int8, int16, int32, int64 - bits
	//Quando usamos só "int" ele considera a arquitetura do nosso computador (32 ou 64 bits)
	var inteiro int8 = 20
	fmt.Println(inteiro)

	var intSystem int = 1000000000000000000 //int32
	fmt.Println(intSystem)

	//uint - São numeros inteiros sem finais
	//var n1 uint = -20 - Error
	var n1 uint = 20
	fmt.Println(n1)

	//Rune = int32
	var n2 rune = 10
	fmt.Println(n2)

	//Byte = int8

	var n3 byte = 10
	fmt.Println(n3)

	/* Float */
	var numeroReal32 float32 = 100.32
	fmt.Println(numeroReal32)

	var numeroReal64 float64 = 10000000.52
	fmt.Println(numeroReal64)

	numeroRealImplicito := 2000.20 //Considera a arquitetura do computador (64 ou 32 bits)
	fmt.Println(numeroRealImplicito)

	/* Strings */
	//Os textos são determinados por "", onde NÃO EXISTE o Char, pois quando informamos o valor de um caracter em 'A' ele devolve
	//o valor inteiro desse caracter na tabela ASC

	var str string = "Leandro"
	str2 := "Souza"

	fmt.Println(str, str2)

	char := 'A' //Output: 65
	fmt.Println(char)

	/* Boolean - Bool */

	var boolean1 bool = true
	boolean2 := false
	fmt.Println(boolean1, boolean2)

}
