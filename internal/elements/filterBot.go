package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildAspectRatio() *widget.Select {
	entries := []string{
		"Any",
		"All Wide",
		"All Portrait",
		"16×9",
		"21×9",
		"9×16",
		"1×1",
		"16×10",
		"32×9",
		"10×16",
		"3×2",
		"48×9",
		"9×18",
		"4×3",
		"5×4",
	}

	aspectRatio := widget.NewSelect(entries, func(s string) {
		log.Printf("Aspect Ratio: %v\n", s)
	})

	return aspectRatio
}

func buildResolution() *widget.Select {
	entries := []string{
		"At Least",
		"Exact",
	}

	resolution := widget.NewSelect(entries, func(s string) {
		log.Printf("Resolution: %v\n", s)
	})

	return resolution
}

func BuildFilterRowBot() *fyne.Container {
	aspectRatioForm := formItem("Ratio", buildAspectRatio())
	resolutionForm := formItem("Resolution", buildResolution())

	filterRow := container.NewGridWithColumns(2, aspectRatioForm, resolutionForm)

	return filterRow
}
