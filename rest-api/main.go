package main

import (
	"fmt"
	"tfalc/GolangStudies/rest-api/models"
	"tfalc/GolangStudies/rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Scarlet Johansson", Historia: "Atriz da Marvel"},
		{Nome: "Chad Bosewick", Historia: "Falecido ator da Marvel"},
	}

	fmt.Println("Iniciando servidor Rest com Golang!")
	routes.HandleRequest()
}
