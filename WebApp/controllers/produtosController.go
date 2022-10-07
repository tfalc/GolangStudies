package controllers

import (
	"net/http"
	"text/template"
	"tfalc/GolangStudies/WebApp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosProdutos := models.BuscaProdutos()

	temp.ExecuteTemplate(w, "Index", todosProdutos)

}
