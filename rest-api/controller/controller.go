package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tfalc/GolangStudies/rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ShowAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
