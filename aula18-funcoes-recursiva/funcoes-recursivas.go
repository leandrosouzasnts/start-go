package main

import "fmt"

func fibonaci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonaci(posicao-2) + fibonaci(posicao-1)
}

func main() {

	fmt.Println(fibonaci(10))
}
