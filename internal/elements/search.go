package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/alunya-go/internal/helper"
	"github.com/Yustinia/gopaper"
)

func BuildSearchSection(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label) *fyne.Container {
	queryEntry := widget.NewEntry()
	queryEntry.SetPlaceHolder("This is your search")
	queryForm := widget.NewForm(
		widget.NewFormItem("Search", queryEntry),
	)
	querySearchButton := widget.NewButton("Enter", func() {
		params.KeySearch = queryEntry.Text
		params.Page = 1

		search.PerformSearch(func() (gopaper.SearchResponse, error) {
			return client.Search(*params)
		}, updateGallery, lastResults, pageLabel)
	})

	searchContainer := container.NewBorder(nil, nil, nil, querySearchButton, queryForm)

	return searchContainer
}
