package main

import "fmt"

func closure() func() {
	texto := "Texto da func closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}
func main() {
	texto := "Texto dentro do main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}
