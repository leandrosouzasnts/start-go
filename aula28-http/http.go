package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loading users ..."))
}

func rotas() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))
	rotas()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
