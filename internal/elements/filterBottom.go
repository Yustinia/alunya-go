package elements

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

var labelToValueRatio = map[string]string{
	"Any":          "",
	"All Wide":     "landscape",
	"All Portrait": "portrait",
	"16x9":         "16x9",
	"21x9":         "21x9",
	"9x16":         "9x16",
	"1x1":          "1x1",
	"16x10":        "16x10",
	"32x9":         "32x9",
	"10x16":        "10x16",
	"3x2":          "3x2",
	"48x9":         "48x9",
	"9x18":         "9x18",
	"4x3":          "4x3",
	"5x4":          "5x4",
}

func buildAspectRatio(params *gopaper.SearchParams) *widget.Select {
	entries := []string{
		"Any",
		"All Wide",
		"All Portrait",
		"16x9",
		"21x9",
		"9x16",
		"1x1",
		"16x10",
		"32x9",
		"10x16",
		"3x2",
		"48x9",
		"9x18",
		"4x3",
		"5x4",
	}

	aspectRatio := widget.NewSelect(entries, func(s string) {
		params.Ratios = labelToValueRatio[s]
	})
	aspectRatio.SetSelected("Any")

	return aspectRatio
}

func buildResolution(params *gopaper.SearchParams) *fyne.Container {
	entries := []string{
		"At Least",
		"Exact",
	}

	resOnX := "1920"
	resOnY := "1080"
	resolutionState := "At Least"

	applyResolution := func() {
		switch resolutionState {
		case "At Least":
			params.AtLeast = fmt.Sprintf("%sx%s", resOnX, resOnY)
			params.Resolution = ""
		case "Exact":
			params.Resolution = fmt.Sprintf("%sx%s", resOnX, resOnY)
			params.AtLeast = ""
		}
	}

	axisXEntry := widget.NewEntry()
	axisXEntry.SetPlaceHolder("1920")
	axisXEntry.SetText(resOnX)
	axisXEntry.OnChanged = func(text string) {
		resOnX = text
		applyResolution()
	}

	axisYEntry := widget.NewEntry()
	axisYEntry.SetPlaceHolder("1080")
	axisYEntry.SetText(resOnY)

	axisYEntry.OnChanged = func(text string) {
		resOnY = text
		applyResolution()
	}

	resolution := widget.NewSelect(entries, func(s string) {
		resolutionState = s
		applyResolution()
	})
	resolution.SetSelected("At Least")

	resolutionContainer := container.NewGridWithColumns(3, resolution, axisXEntry, axisYEntry)

	return resolutionContainer
}

func BuildFilterRowBot(params *gopaper.SearchParams) *fyne.Container {
	aspectRatioForm := formItem("Ratio", buildAspectRatio(params))
	resolutionForm := formItem("Resolution", buildResolution(params))

	filterRow := container.NewGridWithColumns(2, aspectRatioForm, resolutionForm)

	return filterRow
}
