package main

import "fmt"

func main() {

	numero := 10

	if numero > 10 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	if outroNumero := numero; outroNumero > 5 { //if init (porém essa variável (outroNumero) só é válida nesse scopo)
		fmt.Println("Maior que 10")
	}
}
