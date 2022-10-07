package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"tfalc/GolangStudies/WebApp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosProdutos := models.BuscaProdutos()

	temp.ExecuteTemplate(w, "Index", todosProdutos)

}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CriarNovoProduto(nome, descricao, precoFloat, quantidadeInt)
	}

	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	produto := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "EditProduct", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID: ", err.Error())
		}

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err.Error())
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err.Error())
		}

		models.UpdateProduct(idInt, nome, descricao, precoFloat, quantidadeInt)
	}

	http.Redirect(w, r, "/", 301)
}
