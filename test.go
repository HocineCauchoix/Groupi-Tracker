package groupi

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	fynex "fyne.io/x/fyne/widget"
)

// 2008 : 1, 3, 4
// 4: 2, 1,1

func NumberOfMember(artistes []Artistes) {
	var max int
	for i := 0; i < len(artistes); i++ {
		if len(artistes[i].Art.Member) > max {
			max = len(artistes[i].Art.Member)
		}
	}
	fmt.Println(max)
}

func CreationDate(artistes []Artistes) {
	var min int = 2000
	for i := 0; i < len(artistes); i++ {
		if artistes[i].Art.CreationDate < min {
			min = artistes[i].Art.CreationDate
		}
	}
	fmt.Println(min)
}

func PremierAlbum(artistes []Artistes) {
	var min int = 0
	date := ""
	for i := 0; i < len(artistes); i++ {
		date = Année(artistes[i].Art.FirstAlbum)
		nbr, _ := strconv.Atoi(date)
		if nbr > min {
			min = artistes[i].Art.CreationDate
		}
	}
	fmt.Println(min)
}

func Gpt(art []Artistes) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Example")

	truc := []string{}
	entry := fynex.NewCompletionEntry(truc)

	// When the use typed text, complete the list.
	entry.OnChanged = func(s string) {
		// completion start for text length >= 3
		if len(s) < 0 {
			entry.HideCompletion()
			return
		}

		// Make a search on wikipedia

		// Get the list of possible completion
		var results []string
		for i := 0; i < 10; i++ {
			results = append(results, "resultats")
		}

		// no results
		if len(results) == 0 {
			entry.HideCompletion()
			return
		}

		// then show them
		entry.SetOptions(results)
		entry.ShowCompletion()
	}

	// Créer une liste avec quelques éléments
	myList := widget.NewList(
		func() int {
			return 10
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("Element %d", id))
		},
	)

	// Ajouter la liste à un conteneur de défilement
	scrollContainer := container.NewScroll(myList)

	// Créer un conteneur pour afficher les données d'un élément sélectionné
	displayContainer := container.NewVBox()

	// Écouter les événements de sélection de liste
	myList.OnSelected = func(id widget.ListItemID) {
		// Obtenir les données de l'élément sélectionné
		selectedData := fmt.Sprintf("Données de l'élément %d", id)

		// Effacer le contenu précédent du conteneur d'affichage
		displayContainer.Objects = []fyne.CanvasObject{}

		// Ajouter les données de l'élément sélectionné au conteneur d'affichage
		displayContainer.Add(widget.NewLabel(selectedData))
		displayContainer.Add(entry)
	}

	// Créer un conteneur principal pour tous les éléments
	mainContainer := container.NewHSplit(scrollContainer, displayContainer)

	// Définir la taille de la fenêtre
	myWindow.Resize(fyne.NewSize(400, 300))

	// Ajouter le conteneur principal à la fenêtre et afficher la fenêtre
	myWindow.SetContent(mainContainer)
	myWindow.ShowAndRun()
}

func Autocompletion(art []Artistes) {
	myApp := app.New()
	myWindow := myApp.NewWindow("My Search Bar")
	myWindow.Resize(fyne.NewSize(1200, 600))

	completionEntry := fynex.NewCompletionEntry([]string{"apple", "banana", "cherry"})
	completionEntry.SetPlaceHolder("Search...")
	//completionEntry.Wrapping = fyne.TextWrapWord

	// set the max height of the completion dropdown to the height of its container
	// set the max height of the completion dropdown to the height of its container
	// completionEntry.WrappingSetMaxHeight(func() float32 {
	// 	return completionEntry.MinSize().Height
	// })

	// container := fyne.NewContainerWithLayout(
	// 	fynex.NewVBoxLayout(),
	// 	completionEntry,
	// )

	myWindow.SetContent(completionEntry)
	myWindow.ShowAndRun()
}

func LenLoca(tab []Artistes) {
	longueur := 0
	for i := 0; i<len(tab); i++ {
		if len(tab[i].Location) > longueur {
			longueur = len(tab[i].Location)
		}
	}
	fmt.Println(longueur)
}
