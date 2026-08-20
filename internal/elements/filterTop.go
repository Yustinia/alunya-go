package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

func buildCategorySection(params *gopaper.SearchParams) *fyne.Container {
	generalToggle := widget.NewCheck("General", func(b bool) {
		params.Categories.General = b
	})
	generalToggle.SetChecked(params.Categories.General)

	animeToggle := widget.NewCheck("Anime", func(b bool) {
		params.Categories.Anime = b
	})
	animeToggle.SetChecked(params.Categories.Anime)

	peopleToggle := widget.NewCheck("People", func(b bool) {
		params.Categories.People = b
	})
	peopleToggle.SetChecked(params.Categories.People)

	categoryContainer := container.NewGridWithColumns(3, generalToggle, animeToggle, peopleToggle)

	return categoryContainer
}

func buildPuritySection(params *gopaper.SearchParams) *fyne.Container {
	sfwToggle := widget.NewCheck("SFW", func(b bool) {
		params.Purity.SFW = b
	})
	sfwToggle.SetChecked(params.Purity.SFW)

	sketchyToggle := widget.NewCheck("Sketchy", func(b bool) {
		params.Purity.Sketchy = b
	})
	sketchyToggle.SetChecked(params.Purity.Sketchy)

	nsfwToggle := widget.NewCheck("NSFW", func(b bool) {
		params.Purity.NSFW = b
	})
	nsfwToggle.SetChecked(params.Purity.NSFW)

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

func BuildFilterRowTop(params *gopaper.SearchParams) *fyne.Container {
	categoryForm := formItem("Category", buildCategorySection(params))
	purityForm := formItem("Purity", buildPuritySection(params))
	sortForm := formItem("Sort", buildSortSection())

	filterRow := container.NewGridWithColumns(3, categoryForm, purityForm, sortForm)

	return filterRow
}
