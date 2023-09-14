package groupi

import (
	"strings"
)

// Rechercher une artiste
func SearchForArtist(tab []Artistes, Nom string) []Artistes {
	var Art_filte []Artistes

	// Recherche
	for i := 0; i < len(tab); i++ {
		if strings.EqualFold(tab[i].Art.Name, Nom) {
			Art_filte = append(Art_filte, tab[i])
		} else if strings.Contains(strings.ToLower(tab[i].Art.Name), strings.ToLower(Nom)) {
			Art_filte = append(Art_filte, tab[i])
		}
	}
	//fmt.Println(len(Art_filte))

	return Art_filte
}

// Rechercher une date de concert ---------------------------------------
func SearchForCity(tab []Artistes, city string) []Artistes {
	var SearchCity []Artistes
	var boolean bool

	// Recherche
	for i := 0; i < len(tab); i++ {
		boolean = false
		for j := 0; j < len(tab[i].Location); j++ {
			if strings.EqualFold(tab[i].Location[j], city) && !boolean{
				boolean = true
				SearchCity = append(SearchCity, tab[i])
			} else if strings.Contains(strings.ToLower(tab[i].Location[j]), strings.ToLower(city)) && !boolean{
				boolean = true
				SearchCity = append(SearchCity, tab[i])
			}
		}
	}
	return SearchCity
}
