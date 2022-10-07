package routes

import (
	"net/http"
	"tfalc/GolangStudies/WebApp/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/update", controllers.Update)
}
