package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildPageRow() *fyne.Container {
	prevBtn := widget.NewButton("Prev", func() {
		log.Println("Previous")
	})
	nextBtn := widget.NewButton("Next", func() {
		log.Println("Next")
	})
	pageNum := widget.NewEntry()
	pageNum.SetPlaceHolder("Page")
	pageNum.OnChanged = func(s string) {
		log.Printf("Page: %v", s)
	}

	grid := container.NewGridWithColumns(3, prevBtn, pageNum, nextBtn)

	return container.NewCenter(grid)
}
