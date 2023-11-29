package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array")

	var array1 [3]string
	array1[0] = "Leandro"

	array2 := [3]int{1, 2, 3}
	fmt.Println(array1)
	fmt.Println(array2)

	//Mais comum de ser utilizado
	fmt.Println("Slice")

	var slice1 []int

	slice2 := []string{"Leandro", "Divani"}
	fmt.Println(slice1, slice2)

	fmt.Println(reflect.TypeOf(slice1))

	slice1 = append(slice1, 10) //Sempre retorna um novo slice, mas por praticidade podemos usar a mesma variável
	fmt.Println(slice1)

	slice3 := array2[1:3] //Slice é uma fatia de array, por isso é utilizado também como um ponteiro de um determinado array
	fmt.Println(slice3)

	array2[0] = 2
	fmt.Println(slice3)

	//Mesmo foi dobrando de acordo com nossa necessidade, porém ele é criado a partir de um array interno onde definimos o seu respectivo tamanho
	slice4 := make([]int, 2, 3)
	fmt.Println(slice4)
	slice4 = append(slice4, 3)
	fmt.Println(slice4)
	slice4 = append(slice4, 1)
	fmt.Println(slice4)

	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
