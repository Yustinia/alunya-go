package search

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

func PerformSearch(action func() (gopaper.SearchResponse, error), updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label) {
	go func() {
		result, err := action()
		if err != nil {
			log.Println(err)
			return
		}

		fyne.Do(func() {
			updateGallery(result.Wallpapers)
			*lastResults = result
			pageLabel.SetText(strconv.Itoa(result.Metadata.CurrentPage))
		})
	}()
}
