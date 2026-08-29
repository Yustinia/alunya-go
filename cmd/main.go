package main

import (
	"os"
	"strconv"
	"time"

	"github.com/Yustinia/alunya-go/internal/elements"
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

func main() {
	a := app.NewWithID("io.github.yustinia.alunya")
	w := a.NewWindow("alunya")
	winSize := fyne.NewSize(800, 600)
	var lastResults gopaper.SearchResponse
	var debounceTimer *time.Timer
	params := gopaper.NewSearch()
	pageLabel := widget.NewLabel(strconv.Itoa(params.Page))
	pageLabel.Alignment = fyne.TextAlignCenter

	galleryGrid, updateGallery := elements.GalleryGrid()
	searchContainer := elements.BuildSearchSection(&params, &CLIENT, updateGallery, &lastResults, pageLabel, &debounceTimer)
	filterRowTop := elements.BuildFilterRowTop(&params, &CLIENT, updateGallery, &lastResults, pageLabel, &debounceTimer)
	filterRowBot := elements.BuildFilterRowBot(&params)
	pageRow := elements.BuildPageRow(&params, &CLIENT, updateGallery, &lastResults, pageLabel)

	topContainer := container.NewVBox(searchContainer, filterRowTop, filterRowBot)
	windowContainer := container.NewBorder(topContainer, pageRow, nil, nil, galleryGrid)

	w.SetContent(windowContainer)
	w.Resize(winSize)
	w.ShowAndRun()
}
