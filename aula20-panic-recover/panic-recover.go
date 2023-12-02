package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando execução")
	}
}

func media(n1, n2 uint) bool {

	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Média é exatamento 6")
}

func main() {
	fmt.Println("Iniciando programa!")
	fmt.Println(media(10, 2))
	fmt.Println("Finalizando programa!")
}
