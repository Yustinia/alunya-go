package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GalleryGrid() *container.Scroll {
	cellSize := fyne.NewSize(250, 250)
	image := canvas.NewImageFromFile("/home/yustinia/Projects/apps/alunya-go/image/placeholder.jpg")
	image.FillMode = canvas.ImageFillContain
	grid := container.NewGridWrap(cellSize, image)
	scroll := container.NewScroll(grid)

	return scroll
}
