package main

import "fmt"

//Funções com retorno nomeados
func calcMath(n1, n2 int) (sum, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	fmt.Println(calcMath(1, 2))
}
