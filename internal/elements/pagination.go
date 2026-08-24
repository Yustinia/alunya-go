package elements

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

func BuildPageRow(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse) *fyne.Container {
	prevBtn := widget.NewButton("Prev", func() {
		result, err := client.PrevPage(*lastResults, params)
		if err != nil {
			log.Println(err)
			return
		}
		updateGallery(result.Wallpapers)
		*lastResults = result
	})

	nextBtn := widget.NewButton("Next", func() {
		result, err := client.NextPage(*lastResults, params)
		if err != nil {
			log.Println(err)
			return
		}
		updateGallery(result.Wallpapers)
		*lastResults = result
	})

	pageNum := widget.NewEntry()
	pageNum.SetPlaceHolder("Page")
	pageNum.OnSubmitted = func(s string) {
		pageInt, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
			return
		}

		result, err := client.SetPage(*lastResults, params, pageInt)
		if err != nil {
			log.Println(err)
			return
		}
		updateGallery(result.Wallpapers)
		*lastResults = result
	}

	grid := container.NewGridWithColumns(3, prevBtn, pageNum, nextBtn)

	return container.NewCenter(grid)
}
