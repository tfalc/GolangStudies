package routes

import (
	"log"
	"net/http"

	"tfalc/GolangStudies/rest-api/controller"

	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()

	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.ShowAllPersonalities)
	r.HandleFunc("/api/personalidades/{id}", controller.ShowOnePersonality).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
