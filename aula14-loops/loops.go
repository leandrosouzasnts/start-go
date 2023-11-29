package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0

	for i < 10 { //Parecido com While
		i++
		println("Incrementando I")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 5; j++ {
		println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	//Maneira de iterar sobre Array/Slice/Maps/Strings
	nomes := [3]string{"Leandro", "Rafaela", "Divani"}

	for i, valor := range nomes {
		fmt.Println(i, valor)
	}

	user1 := map[string]string{
		"nome":  "Leandro",
		"senha": "123",
	}

	for _, value := range user1 { //Quando não desejo algum valor, então passo " _ "
		fmt.Println(value)
	}

	for key, value := range user1 {
		fmt.Println(key, value)
	}

	for indice, letra := range "TEXTO" {
		//fmt.Println(indice, letra) //O range nos retorna os valores de cada caracter de acordo com a tabela ASCII
		fmt.Println(indice, string(letra))
	}

	//LOOP INFINITO
	// for {

	// }
}
