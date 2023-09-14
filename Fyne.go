package groupi

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"net/http"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	fynex "fyne.io/x/fyne/widget"
)

func Application(artists []Artistes) {

	// Variables --------------------------------------------------------------------------------------------------------------------
	var vide bool = false
	var searchbarON bool = false
	var FiltresOn bool = false
	//test := fynex.NewCompletionEntry()
	var listArtist *widget.List

	ArtSearchName := []Artistes{}
	var Art_filte []Artistes

	var creation float64 = 1957
	var Checkbox [9]bool
	var premier_album float64 = 1961

	// Application --------------------------------------------------------------------------------------------------------------------------
	myApp := app.New()

	// Définir le thème de l'application avec une couleur de fond gris foncé
	myApp.Settings().SetTheme(theme.DarkTheme())
	// Nom Fenêtre
	myWindow := myApp.NewWindow("Napster")
	// Resize la fenêtre
	myWindow.Resize(fyne.NewSize(1300, 700))
	// Création d'un objet image
	logo := canvas.NewImageFromFile("img/Napster.png")
	// Resize image
	logo.Resize(fyne.NewSize(100, 50))
	// Opacité
	//logo.(*canvas.Image).SetOpacity(0.5)

	// Charger l'icône depuis un fichier
	icon, err := fyne.LoadResourceFromPath("img/NapsterIcon.jpg")
	if err != nil {
		fmt.Println("L'icone n'a pas pu être chargée")
	}

	// Crée un conteneur pour l'image
	//imgContainer := container.NewCenter(img)
	// Colors
	gold := color.NRGBA{R: 231, G: 208, B: 77, A: 255}
	// Titre
	title := canvas.NewText("                                                                                                         Napster", gold)

	title.TextStyle = fyne.TextStyle{Italic: true}

	// Création de la barre de recherche + un boutton pour l'activer---------------------------------------------------
	liste_recherche := []string{}
	searchEntry := fynex.NewCompletionEntry(liste_recherche)
	searchEntry.SetPlaceHolder("Rechercher un Artiste, une Ville ou un Pays ...")

	searchButton := widget.NewButton("Rechercher", func() {
		ArtSearchName = SearchForArtist(artists, searchEntry.Text)
		if len(ArtSearchName) == 0 {
			ArtSearchName = SearchForCity(artists, searchEntry.Text)
		}
		searchbarON = true
		FiltresOn = false
		vide = false
		listArtist.Refresh()
	},
	)
	menuItem := &fyne.Menu{
		Label: "File",
		Items: nil,
	}
	menu:= fyne.NewMainMenu(menuItem)

	searchEntry.OnChanged = func(query string) {
		searchbarON = false
		FiltresOn = false
		vide = false
		listArtist.UnselectAll()

		// completion start for text length >= 3
		if len(query) < 2 {
			searchEntry.HideCompletion()
			return
		}

		// Get the list of possible completion
		resultats_art := SearchForArtist(artists, searchEntry.Text)
		resultats_city := SearchForCity(artists, searchEntry.Text)

		resultats := []string{}
		// no result
		if len(resultats_art) == 0 && len(resultats_city) == 0 {
			searchEntry.HideCompletion()
			return
		} else {

			// artiste name
			for i := 0; i < len(resultats_art); i++ {
				resultats = append(resultats, resultats_art[i].Art.Name)
			}

			// ville et pays
			for i := 0; i < len(resultats_city); i++ {
				for j := 0; j < len(resultats_city[i].Location); j++ {
					if strings.Contains(strings.ToLower(resultats_city[i].Location[j]), strings.ToLower(query)) {
						resultats = append(resultats, resultats_city[i].Location[j])
					}
				}
			}
		}

		resultats = PopDoublon(resultats)
		// then show them
		searchEntry.SetOptions(resultats)
		searchEntry.ShowCompletion()
	}
	// Créer un conteneur pour la barre de recherche
	searchContainer := container.NewVBox(searchEntry, searchButton)

	// Creation des Filtres + Bouton pour les appliquer -----------------------------------------------------------------------------------

	// Création du filtre pour la date de création
	creationLabel := widget.NewLabel("Date de création du groupe : ")
	creationValue := widget.NewLabel("")
	creationRange := widget.NewSlider(1957, 2015)
	creationRange.OnChanged = func(value float64) {
		//creationLabel.SetText(fmt.Sprint("Création : "))
		creationValue.SetText(strconv.Itoa(int(value)))
		creation = value
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	}

	// Création des filtre pour le nombre de membres (Check Box)
	memberLabel := widget.NewLabel("Nombre de Membre dans le groupe: ")
	memberValue := widget.NewLabel("")
	un := widget.NewCheck("1 ", func(b bool) {
		Checkbox[1] = b
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	})
	deux := widget.NewCheck("2 ", func(b bool) {
		Checkbox[2] = b
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	})
	trois := widget.NewCheck("3 ", func(b bool) {
		Checkbox[3] = b
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	})
	quatre := widget.NewCheck("4 ", func(b bool) {
		Checkbox[4] = b
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	})
	cinq := widget.NewCheck("5 ", func(b bool) {
		Checkbox[5] = b
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	})
	six := widget.NewCheck("6 ", func(b bool) {
		Checkbox[6] = b
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	})
	sept := widget.NewCheck("7 ", func(b bool) {
		Checkbox[7] = b
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	})
	huit := widget.NewCheck("8 ", func(b bool) {
		Checkbox[8] = b
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	})

	// Création du filtre pour la date du premier album
	albumLabel := widget.NewLabel("Date du premier album du groupe: ")
	albumValue := widget.NewLabel("")
	albumRange := widget.NewSlider(1961, 2018)
	albumRange.OnChanged = func(value float64) {
		premier_album = value
		albumValue.SetText(strconv.Itoa(int(value)))
		searchbarON = false
		FiltresOn = false
		Art_filte = []Artistes{}
		vide = false
		listArtist.UnselectAll()
		listArtist.Refresh()
	}

	// Boutton qui permet d'appliquer les filtres
	// Faire un if pour voir si un des filtres
	appliquer := widget.NewButton("Appliquer", func() {
		var wait_Art []Artistes
		var pas_possible bool = false
		boolean := 0
		FiltresOn = true
		var filtres bool = false
		var check bool = false

		for i := 0; i < len(Checkbox); i++ {
			if Checkbox[i] {
				check = true
			}
		}

		if creation != 1957 || premier_album != 1961 || check {
			filtres = true
		}

		// On vérifie qu'au moins un filtre ait été sélectionner
		if filtres {
			// Checkbox
			for i := 1; i < len(Checkbox); i++ {
				if Checkbox[i] {
					boolean++
					wait_Art = FiltreNombreDeMembre(artists, i)
					if len(wait_Art) == 0 {
						pas_possible = true
					} else {
						for j := 0; j < len(wait_Art); j++ {
							Art_filte = append(Art_filte, wait_Art[j])
						}
					}
				}
			}

			// Slider Creation du groupe
			if !pas_possible && creation != 1957 && boolean != 0 { // Le premier filtre a été activer et c'est possible
				Art_filte = FilterCreationDate(Art_filte, creation)
				if len(Art_filte) == 0 {
					pas_possible = true
				}
			} else if boolean == 0 && creation != 1957 { // Le premier filtre n'a pas été activer
				Art_filte = FilterCreationDate(artists, creation)
				if len(Art_filte) == 0 {
					pas_possible = true
				}
			}

			// Slider Date du premier album
			if !pas_possible && (creation != 1957 || boolean != 0) && premier_album != 1961 { // L'un des deux premier filtre (au moins) a/ont été activer et c'est possible
				Art_filte = FiltrePremierAlbum(Art_filte, premier_album)
				if len(Art_filte) == 0 {
					pas_possible = true
				}
			} else if boolean == 0 && creation == 1957 && premier_album != 1961 { // Aucun des deux autres filtres a été activer
				Art_filte = FiltrePremierAlbum(artists, premier_album)
				if len(Art_filte) == 0 {
					pas_possible = true
				}
			}

			if pas_possible {
				Art_filte = []Artistes{}
			}
			listArtist.UnselectAll()
		} else {
			FiltresOn = false
		}
	})

	// Conteneur des filtres
	filtres := widget.NewLabel("Filtres :")

	premiere_ligne := container.NewHBox(
		creationLabel, creationRange, creationValue,
		albumLabel, albumRange, albumValue,
	)

	deuxieme_ligne := container.NewHBox(
		memberLabel, un, deux, trois, quatre, cinq, six, sept, huit, memberValue,
		appliquer,
	)
	// Layout des widgets (conteneur des filtres)
	content := container.NewVBox(
		filtres,
		premiere_ligne,
		deuxieme_ligne,
	)

	// Liste des Artistes -----------------------------------------------------------------------------------
	listArtist = widget.NewList(func() int {
		if searchbarON {
			if len(ArtSearchName) == 0 {
				vide = true
				return 1
			}
			return len(ArtSearchName)
		} else if FiltresOn {
			if len(Art_filte) == 0 {
				vide = true
				return 1
			}
			return len(Art_filte)
		} else {
			if len(artists) == 0 {
				vide = true
				return 1
			}
			return len(artists)
		}
	}, func() fyne.CanvasObject {
		return widget.NewLabel("musics")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		if searchbarON {
			if vide {
				object.(*widget.Label).SetText("Aucun resultat trouvé ...")
			} else {
				object.(*widget.Label).SetText(ArtSearchName[id].Art.Name)
			}
		} else if FiltresOn {
			if vide {
				object.(*widget.Label).SetText("Aucun resultat trouvé ...")
			} else {
				object.(*widget.Label).SetText(Art_filte[id].Art.Name) //+ "      " + strconv.Itoa(len(Art_filte[id].Art.Member)))
			}
		} else {
			if vide {
				object.(*widget.Label).SetText("Aucun resultat trouvé ...")
			} else {
				object.(*widget.Label).SetText(artists[id].Art.Name)
			}
		}
	})

	// Affichage Info Artiste selectionner ----------------------------------------------------------------------------

	// conteneur des infos
	DateEtLieux := fyne.NewContainerWithLayout(layout.NewGridLayout(4))
	concertContainer := container.NewVBox()
	infoContainer := container.NewVBox()

	// Affichage
	listArtist.OnSelected = func(id widget.ListItemID) {
		var tableau []Artistes
		if searchbarON {
			tableau = ArtSearchName
		} else if FiltresOn {
			tableau = Art_filte
		} else {
			tableau = artists
		}

		// Image --------------------------
		//fmt.Println(tableau[id].Art.Image)
		var contentImg *canvas.Image
		resp, err := http.Get(tableau[id].Art.Image)
		if err != nil {
			// Gestion de l'erreur
			fmt.Println("il y a une erreur")
		}
		defer resp.Body.Close()

		img, err := jpeg.Decode(resp.Body)
		if err != nil {
			// Gestion de l'erreur
			fmt.Println("il y a une erreur")
		}
		//r, _ := fyne.LoadResourceFromURLString()
		contentImg = canvas.NewImageFromImage(img)
		contentImg.FillMode = canvas.ImageFillContain

		contentName := widget.NewLabel("")
		contentName.Wrapping = fyne.TextWrapWord
		contentName.Text = ("Artiste / Groupe : " + tableau[id].Art.Name)

		contentAlbum := widget.NewLabel("")
		contentAlbum.Wrapping = fyne.TextWrapWord
		contentAlbum.Text = ("Premier album : " + tableau[id].Art.FirstAlbum)

		contentMember := widget.NewLabel("")
		contentMember.Wrapping = fyne.TextWrapWord
		contentMember.Text = "Membres : " + strings.Join(tableau[id].Art.Member, ", ")

		contentCreation := widget.NewLabel("")
		contentCreation.Wrapping = fyne.TextWrapWord
		contentCreation.Text = "Date de création : " + strconv.Itoa(tableau[id].Art.CreationDate)

		// dates et lieux de concert -------------------
		concertContainer.Objects = []fyne.CanvasObject{}
		DateEtLieux.Objects = []fyne.CanvasObject{}
		contentConcert := widget.NewLabel("Dates et Lieux de concert :")
		contentConcert.Wrapping = fyne.TextWrapWord
		concertContainer.Add(contentConcert)

		for i := 0; i < len(tableau[id].Location); i++ {
			// Nom de la ville
			lieux := widget.NewLabel(tableau[id].Location[i] + " :")
			lieux.Wrapping = fyne.TextWrapWord

			// Dates de concert
			dates := ""
			for j := 0; j < len(tableau[id].Date[i]); j++ {
				if j == len(tableau[id].Date[i])-1 {
					dates += tableau[id].Date[i][j]
				} else {
					dates += tableau[id].Date[i][j] + ", "
				}
			}
			date := widget.NewLabel(dates)
			date.Wrapping = fyne.TextWrapWord

			// Append aux conteneurs
			affich := container.NewVBox(lieux, date)
			itemContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), affich)

			// Ajoutez le conteneur de l'élément au conteneur principal
			DateEtLieux.Add(itemContainer)
		}

		concertContainer.Add(DateEtLieux)

		// Reload le container
		infoContainer.Objects = []fyne.CanvasObject{}

		infoContainer.Add(container.NewMax(contentImg))
		infoContainer.Add(container.NewMax(contentName))
		infoContainer.Add(container.NewMax(contentAlbum))
		infoContainer.Add(container.NewMax(contentCreation))
		infoContainer.Add(container.NewMax(contentMember))
		infoContainer.Add(container.NewMax(concertContainer))

	}

	// Squelette principal de l'application ------------------------------------------------------------------
	// Créer un conteneur pour logo + title + bdr
	rightContainer := container.NewVBox(
		title,
		canvas.NewRectangle(color.Transparent),
		content,
		searchContainer,
		container.NewMax(infoContainer),
	)

	// Interface
	split := container.NewHSplit(
		listArtist,
		container.NewMax(rightContainer),
	)
	split.Offset = 0.3

	myWindow.SetMainMenu(menu)
	//Set icon
	myWindow.SetIcon(icon)
	//Set Content
	myWindow.SetContent(split)
	// Afficher la fenêtre et commencer la boucle de l'application
	myWindow.ShowAndRun()
}
