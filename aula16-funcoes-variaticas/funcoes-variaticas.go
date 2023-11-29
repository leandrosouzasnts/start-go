package main

import "fmt"

func calMath(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func print(text string, numbers ...int) { //Limitação está na variatica, onde só podemos ter uma por função
	for _, value := range numbers {
		fmt.Println(text, value)
	}
}

func main() {
	fmt.Println(calMath(1, 2, 3, 3, 3, 4))
	print("Number:", 1, 2, 3, 4, 5)
}
