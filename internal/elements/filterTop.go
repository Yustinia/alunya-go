package elements

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/alunya-go/internal/helper"
	"github.com/Yustinia/gopaper"
)

func buildCategorySection(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label, debounceTimer **time.Timer) *fyne.Container {
	var (
		generalToggle *widget.Check
		animeToggle   *widget.Check
		peopleToggle  *widget.Check
	)

	generalToggle = widget.NewCheck("General", func(b bool) {
		newValue := helper.GuardLastEnabled(params.Categories.Anime, params.Categories.People, b)

		if newValue != b {
			generalToggle.SetChecked(true)
			return
		}

		if newValue == params.Categories.General {
			return
		}

		params.Categories.General = newValue
		helper.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	generalToggle.SetChecked(params.Categories.General)

	animeToggle = widget.NewCheck("Anime", func(b bool) {
		newValue := helper.GuardLastEnabled(params.Categories.General, params.Categories.People, b)

		if newValue != b {
			animeToggle.SetChecked(true)
			return
		}

		if newValue == params.Categories.Anime {
			return
		}

		params.Categories.Anime = newValue
		helper.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	animeToggle.SetChecked(params.Categories.Anime)

	peopleToggle = widget.NewCheck("People", func(b bool) {
		newValue := helper.GuardLastEnabled(params.Categories.Anime, params.Categories.General, b)

		if newValue != b {
			peopleToggle.SetChecked(true)
			return
		}

		if newValue == params.Categories.People {
			return
		}

		params.Categories.People = newValue
		helper.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	peopleToggle.SetChecked(params.Categories.People)

	categoryContainer := container.NewGridWithColumns(3, generalToggle, animeToggle, peopleToggle)

	return categoryContainer
}

func buildPuritySection(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label, debounceTimer **time.Timer) *fyne.Container {
	var (
		sfwToggle     *widget.Check
		sketchyToggle *widget.Check
		nsfwToggle    *widget.Check
	)

	sfwToggle = widget.NewCheck("SFW", func(b bool) {
		newValue := helper.GuardLastEnabled(params.Purity.Sketchy, params.Purity.NSFW, b)

		if newValue != b {
			sfwToggle.SetChecked(true)
			return
		}

		if newValue == params.Purity.SFW {
			return
		}

		params.Purity.SFW = newValue
		helper.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	sfwToggle.SetChecked(params.Purity.SFW)

	sketchyToggle = widget.NewCheck("Sketchy", func(b bool) {
		newValue := helper.GuardLastEnabled(params.Purity.SFW, params.Purity.NSFW, b)

		if newValue != b {
			sketchyToggle.SetChecked(true)
			return
		}

		if newValue == params.Purity.Sketchy {
			return
		}

		params.Purity.Sketchy = newValue
		helper.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	sketchyToggle.SetChecked(params.Purity.Sketchy)

	nsfwToggle = widget.NewCheck("NSFW", func(b bool) {
		newValue := helper.GuardLastEnabled(params.Purity.SFW, params.Purity.Sketchy, b)

		if newValue != b {
			nsfwToggle.SetChecked(true)
			return
		}

		if newValue == params.Purity.NSFW {
			return
		}

		params.Purity.NSFW = newValue
		helper.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	nsfwToggle.SetChecked(params.Purity.NSFW)

	purityContainer := container.NewGridWithColumns(3, sfwToggle, sketchyToggle, nsfwToggle)

	return purityContainer
}

var labelToValueSort = map[string]string{
	"Relevance":  "relevance",
	"Random":     "random",
	"Date Added": "date_added",
	"Views":      "views",
	"Favorites":  "favorites",
	"Toplist":    "toplist",
}

func buildSortSection(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label, debounceTimer **time.Timer) *widget.Select {
	entries := []string{
		"Relevance",
		"Random",
		"Date Added",
		"Views",
		"Favorites",
		"Toplist",
	}

	sortDropDown := widget.NewSelect(entries, func(s string) {
		params.Sorting = labelToValueSort[s]
		helper.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)

	})
	sortDropDown.SetSelected("Date Added") // match NewSearch default sort

	return sortDropDown
}

func BuildFilterRowTop(params *gopaper.SearchParams, client *gopaper.Client, updateGallery func([]gopaper.Wallpaper), lastResults *gopaper.SearchResponse, pageLabel *widget.Label, debounceTimer **time.Timer) *fyne.Container {
	categoryForm := formItem("Category", buildCategorySection(params, client, updateGallery, lastResults, pageLabel, debounceTimer))
	purityForm := formItem("Purity", buildPuritySection(params, client, updateGallery, lastResults, pageLabel, debounceTimer))
	sortForm := formItem("Sort", buildSortSection(params, client, updateGallery, lastResults, pageLabel, debounceTimer))

	filterRow := container.NewGridWithColumns(3, categoryForm, purityForm, sortForm)

	return filterRow
}
