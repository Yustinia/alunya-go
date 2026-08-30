package helper

import (
	"log"
	"strconv"
	"time"

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

func TriggerDebounceSearch(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label, debounceTimer **time.Timer) {
	params.Page = 1

	if *debounceTimer != nil {
		(*debounceTimer).Stop()
	}

	*debounceTimer = time.AfterFunc(800*time.Millisecond, func() {
		PerformSearch(func() (gopaper.SearchResponse, error) {
			return client.Search(*params)
		}, updateGallery, lastResults, pageLabel)
	})
}
