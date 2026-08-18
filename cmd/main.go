package main

import (
	"log"
	"os"

	"github.com/Yustinia/gopaper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	API_KEY = os.Getenv("WALLHAVEN_API_KEY")
	CLIENT  = gopaper.NewClient(API_KEY)
)

// func main() {
// 	params := gopaper.NewSearch()
// 	params.KeySearch = "japan"
//
// 	result, err := CLIENT.Search(params)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//
// 	for i, wall := range result.Wallpapers {
// 		log.Println(i, wall.Path)
// 	}
// }

func buildSearchSection() *fyne.Container {
	queryEntry := widget.NewEntry()
	queryEntry.SetPlaceHolder("This is your search")
	queryForm := widget.NewForm(
		widget.NewFormItem("Search", queryEntry),
	)
	querySearchButton := widget.NewButton("Enter", func() {
		log.Printf("Searched: %v\n", queryEntry.Text)
	})

	searchContainer := container.NewBorder(nil, nil, nil, querySearchButton, queryForm)

	return searchContainer
}

func buildCategorySection() *fyne.Container {
	generalToggle := widget.NewCheck("General", func(b bool) {
		log.Printf("General: %v\n", b)
	})
	animeToggle := widget.NewCheck("Anime", func(b bool) {
		log.Printf("Anime: %v\n", b)
	})
	peopleToggle := widget.NewCheck("People", func(b bool) {
		log.Printf("People: %v\n", b)
	})

	categoryContainer := container.NewGridWithColumns(3, generalToggle, animeToggle, peopleToggle)

	return categoryContainer
}

func main() {
	a := app.NewWithID("alunya")
	w := a.NewWindow("alunya")

	searchContainer := buildSearchSection()
	categoryContainer := buildCategorySection()

	windowContainer := container.NewVBox(searchContainer, categoryContainer)

	w.SetContent(windowContainer)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
