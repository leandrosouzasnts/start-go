package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Nome  string `json:nome`
	Raca  string `json:raca`
	Idade uint   `json:-` //não aparece a tag no json
}

func main() {

	d := dog{"Rex", "Vira-la", 5}

	dogJSON, erro := json.Marshal(d) //Transforma uma struct ou map em JSON

	if erro != nil {
		log.Fatal(erro)
	}

	d1 := map[string]string{
		"nome": "Bolinha",
		"raca": "Pintcher",
	}

	dog1JSON, erro := json.Marshal(d1)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Printf("Struct: %s\n", bytes.NewBuffer(dogJSON))
	fmt.Printf("Map: %s\n", bytes.NewBuffer(dog1JSON))

	//Transformando Json em Struct

	dJSON := `{"Nome": "Rex", "Raca": "Vira-lata", "Idade": 5}`

	var dogStruct dog

	if erro := json.Unmarshal([]byte(dJSON), &dogStruct); erro != nil {
		log.Fatal(erro)
	}

	dogStruct.Raca = "Pastor alemão"
	fmt.Println(dogStruct)
}
