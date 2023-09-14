package groupi

import "strconv"

// Recherche par date de créations du groupe
func FilterCreationDate(tab []Artistes, creation float64) []Artistes {
	var CreationDate_filtre []Artistes
	for i := 0; i < len(tab); i++ {
		if tab[i].Art.CreationDate == int(creation) {
			CreationDate_filtre = append(CreationDate_filtre, tab[i])
		}
	}
	return CreationDate_filtre
}

// Recherche par nombre de membre dans le groupe ---------------------------------------
func FiltreNombreDeMembre(tab []Artistes, nbr int) []Artistes {
	var NumberOfMember []Artistes
	for i := 0; i < len(tab); i++ {
		if len(tab[i].Art.Member) == nbr {
			NumberOfMember = append(NumberOfMember, tab[i])
		}
	}
	return NumberOfMember
}

// Recherche par date de sortie du premiere album ---------------------------------------
func FiltrePremierAlbum(tab []Artistes, premier_album float64) []Artistes {
	var FirstAlbumDate_filtre []Artistes
	var annee string

	for i := 0; i < len(tab); i++ {
		annee = Année(tab[i].Art.FirstAlbum)
		nbr, _ := strconv.Atoi(annee)
		if nbr == int(premier_album) {
			FirstAlbumDate_filtre = append(FirstAlbumDate_filtre, tab[i])
		}
	}

	return FirstAlbumDate_filtre
}

func Année(str string) string {
	str2 := ""
	byte_str := []byte(str)
	for i := len(byte_str) - 1; i > len(byte_str)-5; i-- {
		str2 = string(byte_str[i]) + str2
	}
	return str2
}
