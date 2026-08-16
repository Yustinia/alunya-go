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

func main() {
	params := gopaper.NewSearch()
	params.KeySearch = "japan"

	result, err := CLIENT.Search(params)
	if err != nil {
		log.Fatalln(err)
	}

	for i, wall := range result.Wallpapers {
		log.Println(i, wall.Path)
	}
}
