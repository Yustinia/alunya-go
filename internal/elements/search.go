package elements

import (
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/alunya-go/internal/helper"
	"github.com/Yustinia/gopaper"
)

func BuildSearchSection(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label, debounceTimer **time.Timer) *widget.Form {
	queryEntry := widget.NewEntry()
	queryEntry.SetPlaceHolder("This is your search")
	queryEntry.OnChanged = func(text string) {
		params.KeySearch = text
		helper.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	}

	queryForm := widget.NewForm(
		widget.NewFormItem("Search", queryEntry),
	)

	return queryForm
}
