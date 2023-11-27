package main

import "fmt"

type user struct {
	name     string
	age      uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	cep        string
	bairro     string
	numero     string
}

func main() {

	var u user
	u.name = "Leandro"
	u.age = 26

	//e := endereco{"Santo Antonio", "19000000", "Teste", "400"}
	user1 := user{"Leandro", 26, endereco{"Santo Antonio", "19000000", "Teste", "400"}}

	user2 := user{age: 26}

	fmt.Println(user1, user2)
}
