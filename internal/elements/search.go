package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

func BuildSearchSection(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper)) *fyne.Container {
	queryEntry := widget.NewEntry()
	queryEntry.SetPlaceHolder("This is your search")
	queryForm := widget.NewForm(
		widget.NewFormItem("Search", queryEntry),
	)
	querySearchButton := widget.NewButton("Enter", func() {
		params.KeySearch = queryEntry.Text

		result, err := client.Search(*params)
		if err != nil {
			log.Println(err)
		}
		updateGallery(result.Wallpapers)
	})

	searchContainer := container.NewBorder(nil, nil, nil, querySearchButton, queryForm)

	return searchContainer
}
