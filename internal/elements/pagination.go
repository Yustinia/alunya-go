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
	pageLabel := widget.NewLabel(strconv.Itoa(params.Page))
	pageLabel.Alignment = fyne.TextAlignCenter

	prevBtn := widget.NewButton("Prev", func() {
		go func() {
			result, err := client.PrevPage(*lastResults, params)
			if err != nil {
				log.Println(err)
				return
			}

			fyne.Do(func() {
				updateGallery(result.Wallpapers)
				*lastResults = result

				pageLabel.SetText(strconv.Itoa(lastResults.Metadata.CurrentPage))
			})
		}()
	})

	nextBtn := widget.NewButton("Next", func() {
		go func() {
			result, err := client.NextPage(*lastResults, params)
			if err != nil {
				log.Println(err)
				return
			}

			fyne.Do(func() {
				updateGallery(result.Wallpapers)
				*lastResults = result

				pageLabel.SetText(strconv.Itoa(lastResults.Metadata.CurrentPage))
			})
		}()
	})

	grid := container.NewGridWithColumns(3, prevBtn, pageLabel, nextBtn)

	return container.NewCenter(grid)
}
