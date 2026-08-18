package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildSearchSection() *fyne.Container {
	queryEntry := widget.NewEntry()
	queryEntry.SetPlaceHolder("This is your search")
	queryForm := widget.NewForm(
		widget.NewFormItem("Search", queryEntry),
	)
	querySearchButton := widget.NewButton("Enter", func() {
		log.Printf("Searched: %v\n", queryEntry.Text)
	})

	searchContainer := container.NewBorder(nil, nil, nil, querySearchButton, queryForm)

	return searchContainer
}

func buildCategorySection() *fyne.Container {
	generalToggle := widget.NewCheck("General", func(b bool) {
		log.Printf("General: %v\n", b)
	})
	animeToggle := widget.NewCheck("Anime", func(b bool) {
		log.Printf("Anime: %v\n", b)
	})
	peopleToggle := widget.NewCheck("People", func(b bool) {
		log.Printf("People: %v\n", b)
	})

	categoryContainer := container.NewGridWithColumns(3, generalToggle, animeToggle, peopleToggle)

	return categoryContainer
}

func buildPuritySection() *fyne.Container {
	sfwToggle := widget.NewCheck("SFW", func(b bool) {
		log.Printf("SFW: %v\n", b)
	})
	sketchyToggle := widget.NewCheck("Sketchy", func(b bool) {
		log.Printf("Sketchy: %v\n", b)
	})
	nsfwToggle := widget.NewCheck("NSFW", func(b bool) {
		log.Printf("NSFW: %v\n", b)
	})

	purityContainer := container.NewGridWithColumns(3, sfwToggle, sketchyToggle, nsfwToggle)

	return purityContainer
}

func buildSortSection() *widget.Select {
	entries := []string{
		"Relevance",
		"Random",
		"Date Added",
		"Views",
		"Favorites",
		"Toplist",
		"Hot",
	}
	sortDropDown := widget.NewSelect(entries, func(s string) {
		log.Printf("Sort: %v\n", s)
	})

	return sortDropDown
}

func BuildFilterRow() *fyne.Container {
	categoryForm := formItem("Category", buildCategorySection())
	purityForm := formItem("Purity", buildPuritySection())
	sortForm := formItem("Sort", buildSortSection())

	filterRow := container.NewGridWithColumns(3, categoryForm, purityForm, sortForm)

	return filterRow
}

func formItem(label string, container fyne.CanvasObject) *widget.Form {
	return widget.NewForm(
		widget.NewFormItem(label, container),
	)
}
