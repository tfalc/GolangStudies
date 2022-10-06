package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBanco() *sql.DB {
	conexao := "user=postgres dbname=tfalc_loja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaComBanco()
	defer db.Close()
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
