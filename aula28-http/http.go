package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Seja bem-vindo!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loading users ..."))
}

func rotas() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)
}

func main() {

	rotas()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
