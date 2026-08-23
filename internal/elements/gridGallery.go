package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GalleryGrid() *container.Scroll {
	wallpaperCount := 24
	length := func() int {
		return wallpaperCount
	}

	createImage := func() fyne.CanvasObject {
		image := canvas.NewImageFromFile("image/placeholder.jpg")
		image.FillMode = canvas.ImageFillContain
		image.SetMinSize(fyne.NewSize(250, 250))
		return image
	}

	updateItem := func(id widget.GridWrapItemID, item fyne.CanvasObject) {
	}

	grid := widget.NewGridWrap(length, createImage, updateItem)
	scroll := container.NewScroll(grid)

	return scroll
}
