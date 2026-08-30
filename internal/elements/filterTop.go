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
		if !b && !params.Categories.Anime && !params.Categories.People {
			generalToggle.SetChecked(true)
			return
		}

		if b == params.Categories.General {
			return
		}

		params.Categories.General = b
		search.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	generalToggle.SetChecked(params.Categories.General)

	animeToggle = widget.NewCheck("Anime", func(b bool) {
		if !b && !params.Categories.General && !params.Categories.People {
			animeToggle.SetChecked(true)
			return
		}

		if b == params.Categories.Anime {
			return
		}

		params.Categories.Anime = b
		search.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	animeToggle.SetChecked(params.Categories.Anime)

	peopleToggle = widget.NewCheck("People", func(b bool) {
		if !b && !params.Categories.Anime && !params.Categories.General {
			peopleToggle.SetChecked(true)
			return
		}

		if b == params.Categories.People {
			return
		}

		params.Categories.People = b
		search.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)

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
		if !b && !params.Purity.Sketchy && !params.Purity.NSFW {
			sfwToggle.SetChecked(true)
			return
		}

		if b == params.Purity.SFW {
			return
		}

		params.Purity.SFW = b
		search.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	sfwToggle.SetChecked(params.Purity.SFW)

	sketchyToggle = widget.NewCheck("Sketchy", func(b bool) {
		if !b && !params.Purity.SFW && !params.Purity.NSFW {
			sketchyToggle.SetChecked(true)
			return
		}

		if b == params.Purity.Sketchy {
			return
		}

		params.Purity.Sketchy = b
		search.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
	})
	sketchyToggle.SetChecked(params.Purity.Sketchy)

	nsfwToggle = widget.NewCheck("NSFW", func(b bool) {
		if !b && !params.Purity.SFW && !params.Purity.Sketchy {
			nsfwToggle.SetChecked(true)
			return
		}

		if b == params.Purity.NSFW {
			return
		}

		params.Purity.NSFW = b
		search.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)
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
		search.TriggerDebounceSearch(params, client, updateGallery, lastResults, pageLabel, debounceTimer)

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
