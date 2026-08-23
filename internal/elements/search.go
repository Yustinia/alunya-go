package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

func BuildSearchSection(params *gopaper.SearchParams) *fyne.Container {
	queryEntry := widget.NewEntry()
	queryEntry.SetPlaceHolder("This is your search")
	queryForm := widget.NewForm(
		widget.NewFormItem("Search", queryEntry),
	)
	querySearchButton := widget.NewButton("Enter", func() {
		params.KeySearch = queryEntry.Text
	})

	searchContainer := container.NewBorder(nil, nil, nil, querySearchButton, queryForm)

	return searchContainer
}
