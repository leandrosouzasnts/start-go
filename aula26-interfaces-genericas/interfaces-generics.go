package main

import "fmt"

//Isso é usado para "burlarmos" o controle de tipos, onde devemos tomar cuidado
//Em alguns casos faz sentido usar esse tipo de situação, como é o caso do println

func qualquerCoisa(any interface{}) {
	fmt.Println(any)
}

type User struct {
	name string
}

func main() {

	qualquerCoisa(1)
	qualquerCoisa("Olá")
	qualquerCoisa(true)

	user := User{"Carlos"}
	qualquerCoisa(user)
}
