package main

import "fmt"

type User struct {
	name string
	age  uint
}

func (u User) printInf() { //É uma copia
	fmt.Printf("Nome: %s, Idade: %d \n", u.name, u.age)
}

func (u *User) haveBirthday() {
	u.age++
}

func (u *User) alterName(name string) { //Usar ponteiros para alterar dados
	u.name = name
}

func main() {

	user := User{"Leandro", 26}
	user.printInf()

	user.haveBirthday()
	user.printInf()

	user.alterName("Carlos")
	user.printInf()
}
