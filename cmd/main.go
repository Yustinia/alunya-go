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

func main() {
	a := app.NewWithID("alunya")
	w := a.NewWindow("alunya")

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Enter Search")

	sfw := widget.NewCheck("SFW", func(b bool) {
		log.Println("SFW:", b)
	})
	sketchy := widget.NewCheck("Sketchy", func(b bool) {
		log.Println("Sketchy:", b)
	})
	nsfw := widget.NewCheck("NSFW", func(b bool) {
		log.Println("NSFW:", b)
	})
	purityRow := container.NewHBox(sfw, sketchy, nsfw)

	general := widget.NewCheck("General", func(b bool) {
		log.Println("General:", b)
	})
	anime := widget.NewCheck("Anime", func(b bool) {
		log.Println("Anime:", b)
	})
	people := widget.NewCheck("People", func(b bool) {
		log.Println("People:", b)
	})
	categoryRow := container.NewHBox(general, anime, people)

	sortDropDown := widget.NewSelect([]string{"Relevance", "Date Added", "Random"}, func(s string) {
		log.Println("Sort:", s)
	})

	catPurSortRow := container.NewHBox(categoryRow, purityRow, sortDropDown)

	aspectRatio := widget.NewSelect([]string{"Any", "All Wide", "16:9"}, func(s string) {
		log.Println("Ratio:", s)
	})
	resolution := widget.NewSelect([]string{"At Least", "Exact"}, func(s string) {
		log.Println("Resolution:", s)
	})
	resXAxis := widget.NewEntry()
	resXAxis.SetPlaceHolder("1920")
	resYAxis := widget.NewEntry()
	resYAxis.SetPlaceHolder("1080")

	ratResXYRow := container.NewHBox(aspectRatio, resolution, resXAxis, resYAxis)

	filterSection := container.NewVBox(searchEntry, catPurSortRow, ratResXYRow)

	w.SetContent(filterSection)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
