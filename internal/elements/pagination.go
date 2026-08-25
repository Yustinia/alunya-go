package elements

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

func BuildPageRow(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse) *fyne.Container {
	pageLabel := widget.NewLabel(strconv.Itoa(params.Page))
	pageLabel.Alignment = fyne.TextAlignCenter

	prevBtn := widget.NewButton("Prev", func() {
		result, err := client.PrevPage(*lastResults, params)
		if err != nil {
			log.Println(err)
			return
		}
		updateGallery(result.Wallpapers)
		*lastResults = result

		pageLabel.SetText(fmt.Sprintf("%v/%v", strconv.Itoa(lastResults.Metadata.CurrentPage), result.Metadata.LastPage))
	})

	nextBtn := widget.NewButton("Next", func() {
		result, err := client.NextPage(*lastResults, params)
		if err != nil {
			log.Println(err)
			return
		}
		updateGallery(result.Wallpapers)
		*lastResults = result

		pageLabel.SetText(fmt.Sprintf("%v/%v", strconv.Itoa(lastResults.Metadata.CurrentPage), result.Metadata.LastPage))
	})

	grid := container.NewGridWithColumns(3, prevBtn, pageLabel, nextBtn)

	return container.NewCenter(grid)
}
