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
	thumbCache := make(map[string]fyne.Resource)

	var wallpapers []gopaper.Wallpaper
	length := func() int {
		return len(wallpapers)
	}

	createImage := func() fyne.CanvasObject {
		image := canvas.NewImageFromFile("image/noThumb.jpg")
		image.FillMode = canvas.ImageFillContain
		image.SetMinSize(fyne.NewSize(250, 250))
		return image
	}

	updateItem := func(id widget.GridWrapItemID, item fyne.CanvasObject) {
		url := wallpapers[id].ThumbSmall()
		img := item.(*canvas.Image)

		if value, ok := thumbCache[url]; ok {
			fyne.Do(func() {
				img.Resource = value
				img.Refresh()
			})
			return
		}

		fyne.Do(func() {
			img.Resource = nil
			img.Refresh()
		})

		go func() {
			wall, err := storage.ParseURI(url)
			if err != nil {
				log.Println(err)
				return
			}

			res, err := storage.LoadResourceFromURI(wall)
			if err != nil {
				log.Println(err)
				return
			}

			fyne.Do(func() {
				img.Resource = res
				img.Refresh()
				thumbCache[url] = res
			})
		}()
	}

	grid := widget.NewGridWrap(length, createImage, updateItem)
	scroll := container.NewScroll(grid)

	updateFunc := func(newWallpapers []gopaper.Wallpaper) {
		wallpapers = newWallpapers
		grid.Refresh()
	}

	return scroll, updateFunc
}
