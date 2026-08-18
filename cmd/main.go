package main

import (
	"os"

	"github.com/Yustinia/alunya-go/internal/elements"
	"github.com/Yustinia/gopaper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

func main() {
	a := app.NewWithID("alunya")
	w := a.NewWindow("alunya")

	searchContainer := elements.BuildSearchSection()
	filterRow := elements.BuildFilterRow()

	windowContainer := container.NewVBox(searchContainer, filterRow)

	w.SetContent(windowContainer)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
