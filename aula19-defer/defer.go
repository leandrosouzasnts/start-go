package main

import "fmt"

func funcao1() {
	fmt.Println("funcao 01")
}

func funcao2() {
	fmt.Println("funcao 02")
}

func isAprove(n1, n2 uint) bool {
	fmt.Println("Is aprove")
	defer fmt.Println("Resultado aprovação") //Com defer adiamos até o ultimo momento a execução de um trecho de código.
	if (n1 + n2/2) > 5 {
		return true
	} else {
		return false
	}

}

func main() {

	funcao1()
	funcao2()

	fmt.Println(isAprove(10, 5))
}
