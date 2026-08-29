package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/alunya-go/internal/helper"
	"github.com/Yustinia/gopaper"
)

func BuildPageRow(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label) *fyne.Container {
	prevBtn := widget.NewButton("Prev", func() {
		pageLabel.SetText("Loading...")

		search.PerformSearch(
			func() (gopaper.SearchResponse, error) {
				return client.PrevPage(*lastResults, params)
			}, updateGallery, lastResults, pageLabel)
	})

	nextBtn := widget.NewButton("Next", func() {
		pageLabel.SetText("Loading...")

		search.PerformSearch(func() (gopaper.SearchResponse, error) {
			return client.NextPage(*lastResults, params)
		}, updateGallery, lastResults, pageLabel)
	})

	grid := container.NewGridWithColumns(3, prevBtn, pageLabel, nextBtn)

	return container.NewCenter(grid)
}
