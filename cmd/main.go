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
	a := app.NewWithID("io.github.yustinia.alunya")
	w := a.NewWindow("alunya")

	searchContainer := elements.BuildSearchSection()
	filterRowTop := elements.BuildFilterRowTop()
	filterRowBot := elements.BuildFilterRowBot()
	gridTest := elements.GalleryGrid()

	windowContainer := container.NewVBox(searchContainer, filterRowTop, filterRowBot, gridTest)

	w.SetContent(windowContainer)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
