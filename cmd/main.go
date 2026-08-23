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
	winSize := fyne.NewSize(800, 600)
	params := gopaper.NewSearch()

	galleryGrid, updateGallery := elements.GalleryGrid()
	searchContainer := elements.BuildSearchSection(&params, &CLIENT, updateGallery)
	filterRowTop := elements.BuildFilterRowTop(&params)
	filterRowBot := elements.BuildFilterRowBot(&params)
	pageRow := elements.BuildPageRow()

	topContainer := container.NewVBox(searchContainer, filterRowTop, filterRowBot)
	windowContainer := container.NewBorder(topContainer, pageRow, nil, nil, galleryGrid)

	w.SetContent(windowContainer)
	w.Resize(winSize)
	w.ShowAndRun()
}
