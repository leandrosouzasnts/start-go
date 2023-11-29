package main

import "fmt"

func weekDay(day int) string {
	//case convencional
	switch day {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func weekDay1(day int) string {

	switch {
	case day == 1:
		return "Domingo"
		//fallthrough //joga para a segunda condição
	default:
		return "Valor inválido"
	}
}
func main() {
	fmt.Println(weekDay(1))
	fmt.Println(weekDay1(1))
}
