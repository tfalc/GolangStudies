package routes

import (
	"log"
	"net/http"

	"tfalc/GolangStudies/rest-api/controller"
)

func HandleRequest() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/personalidades", controller.ShowAllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
