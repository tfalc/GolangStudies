package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 39.90, Quantidade: 5},
		{"Tênis", "Preto e azul", 79.90, 15},
		{"Saia", "Rosa", 39.90, 20},
		{"Calça", "Jeans", 89.90, 13},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
