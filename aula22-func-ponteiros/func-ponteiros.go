package main

import "fmt"

func invertSignal(numero *int) { //Recebendo valor do ponteiro

	*numero = *numero * -1
}

func main() {
	numero := -1
	invertSignal(&numero)
	fmt.Println(numero)
}
