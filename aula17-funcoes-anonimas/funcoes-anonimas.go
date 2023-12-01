package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido por parametro: %s", texto)
	}("Hello word")

	fmt.Println(retorno)
}
