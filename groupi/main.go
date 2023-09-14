package main

import (
	"groupi-tracker"
)

func main() {

	// Creation de la structure avec toutes les données
	artistes := groupi.Remplissage()

	// Lancement de l'application
	groupi.Application(artistes)

	// Tests
	//groupi.Gpt(artistes)
	//groupi.Autocompletion(artistes)
	//groupi.LenLoca(artistes)

}
