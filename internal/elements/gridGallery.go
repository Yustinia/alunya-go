package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

func GalleryGrid() (*container.Scroll, func([]gopaper.Wallpaper)) {
	var wallpapers []gopaper.Wallpaper
	length := func() int {
		return len(wallpapers)
	}

	createImage := func() fyne.CanvasObject {
		image := canvas.NewImageFromFile("image/placeholder.jpg")
		image.FillMode = canvas.ImageFillContain
		image.SetMinSize(fyne.NewSize(250, 250))
		return image
	}

	updateItem := func(id widget.GridWrapItemID, item fyne.CanvasObject) {
		wall, err := storage.ParseURI(wallpapers[id].ThumbSmall())
		if err != nil {
			log.Println(err)
			return
		}

		img := item.(*canvas.Image)

		res, err := storage.LoadResourceFromURI(wall)
		if err != nil {
			log.Println(err)
			return
		}

		img.Resource = res
		img.Refresh()
	}

	grid := widget.NewGridWrap(length, createImage, updateItem)
	scroll := container.NewScroll(grid)

	updateFunc := func(newWallpapers []gopaper.Wallpaper) {
		wallpapers = newWallpapers
		grid.Refresh()
	}

	return scroll, updateFunc
}
