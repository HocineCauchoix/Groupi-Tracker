package groupi

import (
	//"encoding/json"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	//"io/ioutil"
	"log"
	"net/http"
)

// On recupere les liens des autres API
func GetDataApi(url string) Apis {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	// on creer la request (ce qu'on veut faire (recuperer des données), on donne l'url, ...)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// on envoie la request
	// on recupere la reponse dans res
	response, getErr := spaceClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// On regarde si le body n'est pas vide (le body de la page)
	if response.Body != nil {
		defer response.Body.Close()
	}

	// On lit le body de la page (on la transforme en tableau de byte)
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Puis on met le body dans une structure qu'on retourne
	var api Apis
	jsonErr := json.Unmarshal(body, &api)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return api
}

// On recupere les donnée de l'API Artiste
func GetDataArtistes(url string) []Art {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	// on creer la request (ce qu'on veut faire (recuperer des données), on donne l'url, ...)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// on envoie la request
	// on recupere la reponse dans res
	response, getErr := spaceClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// On regarde si le body n'est pas vide (le body de la page)
	if response.Body != nil {
		defer response.Body.Close()
	}

	// On lit le body de la page (on la transforme en tableau de byte)
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Puis on met le body dans une structure qu'on retourne
	var artistes []Art
	jsonErr := json.Unmarshal(body, &artistes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//fmt.Println(artistes[0])
	//fmt.Println(len(artistes))
	return artistes
}

// On recupere les donnée de l'API Locations
func GetDataLocations(url string) Location {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	// on creer la request (ce qu'on veut faire (recuperer des données), on donne l'url, ...)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// on envoie la request
	// on recupere la reponse dans res
	response, getErr := spaceClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// On regarde si le body n'est pas vide (le body de la page)
	if response.Body != nil {
		defer response.Body.Close()
	}

	// On lit le body de la page (on la transforme en tableau de byte)
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Puis on met le body dans une structure qu'on retourne
	var loca Location
	jsonErr := json.Unmarshal(body, &loca)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//fmt.Println(rela.Ind[0])
	return loca
}

func GetDataRelations(url string) Relation {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	// on creer la request (ce qu'on veut faire (recuperer des données), on donne l'url, ...)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// on envoie la request
	// on recupere la reponse dans res
	response, getErr := spaceClient.Do(request)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// On regarde si le body n'est pas vide (le body de la page)
	if response.Body != nil {
		defer response.Body.Close()
	}

	// On lit le body de la page (on la transforme en tableau de byte)
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Puis on met le body dans une structure qu'on retourne
	var relations Relation
	jsonErr := json.Unmarshal(body, &relations)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//fmt.Println(rela.Ind[0])
	return relations
}

// Nettoyages des données
func Remplissage() []Artistes {

	// creation variables
	var artistes []Artistes
	var apis Apis
	var art []Art
	var location Location
	var relation Relation
	url_base := "https://groupietrackers.herokuapp.com/api"

	// recuperations des données dans les api
	apis = GetDataApi(url_base)
	art = GetDataArtistes(apis.Artists)
	location = GetDataLocations(apis.Locations)
	relation = GetDataRelations(apis.Relations)

	// fmt.Println(apis.Dates)
	// fmt.Println(apis.Relations)
	// fmt.Println(relation.Index[2])

	// ajout des tableaux date et location à artiste
	for i := 0; i < len(art); i++ {
		var truc Artistes
		truc.Art = &art[i]
		truc.Location = location.Index[i].Location
		for y := 0; y < len(location.Index[i].Location); y++ {
			truc.Date = append(truc.Date, relation.Index[i].Relation[location.Index[i].Location[y]])
		}
		artistes = append(artistes, truc)
	}

	// On enleve les étoiles des dates
	for i := 0; i < len(artistes); i++ {
		for y := 0; y < len(artistes[i].Date); y++ {
			for j := 0; j < len(artistes[i].Date[y]); j++ {
				if len(artistes[i].Date[y][j]) >= 11 {
					artistes[i].Date[y][j] = strings.TrimLeft(artistes[i].Date[y][j], "*")
				}
			}
		}
	}

	// On enleve les tirets et underscores des locations
	for i := 0; i < len(artistes); i++ {
		for y := 0; y < len(artistes[i].Location); y++ {
			var str string
			tab := strings.Split(artistes[i].Location[y], "")
			for j := 0; j < len(tab); j++ {
				if tab[j] == "-" {
					str += " "
				} else if tab[j] == "_" {
					str += " "
				} else {
					str += tab[j]
				}
			}
			artistes[i].Location[y] = str
		}
		//fmt.Println(artistes[i].Location)
	}

	return artistes
}
