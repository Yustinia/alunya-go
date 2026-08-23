package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GalleryGrid() *container.Scroll {
	var images []fyne.CanvasObject
	cellSize := fyne.NewSize(250, 250)

	for i := 0; i < 24; i++ {
		image := canvas.NewImageFromFile("image/placeholder.jpg")
		image.FillMode = canvas.ImageFillContain
		images = append(images, image)
	}
	grid := container.NewGridWrap(cellSize, images...)
	scroll := container.NewScroll(grid)

	return scroll
}
