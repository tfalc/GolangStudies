package main

import (
	"net/http"
	"tfalc/GolangStudies/WebApp/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
