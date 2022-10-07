package routes

import (
	"net/http"
	"tfalc/GolangStudies/WebApp/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NewProduct)
}
