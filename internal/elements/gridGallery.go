package elements

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/Yustinia/gopaper"
)

func GalleryGrid() (*container.Scroll, func([]gopaper.Wallpaper)) {
	green := color.RGBA{G: 255, A: 255}
	red := color.RGBA{R: 255, A: 255}
	yellow := color.RGBA{R: 255, G: 255, A: 255}
	size := fyne.NewSize(341, 192)
	borderSz := 4
	borderRad := 8

	thumbCache := make(map[string]fyne.Resource)

	var wallpapers []gopaper.Wallpaper
	length := func() int {
		return len(wallpapers)
	}

	createImage := func() fyne.CanvasObject {
		image := canvas.NewImageFromFile("image/noThumb.jpg")
		image.CornerRadius = float32(borderRad)

		bg := canvas.NewRectangle(color.Transparent)
		bg.StrokeColor = color.RGBA{
			A: 0,
		}
		bg.StrokeWidth = float32(borderSz)
		bg.CornerRadius = float32(borderRad)

		image.FillMode = canvas.ImageFillCover
		image.SetMinSize(size)
		bg.SetMinSize(size)

		stack := container.NewStack(image, bg)
		return stack
	}

	updateItem := func(id widget.GridWrapItemID, item fyne.CanvasObject) {

		url := wallpapers[id].ThumbSmall()

		stack := item.(*fyne.Container)
		img := stack.Objects[0].(*canvas.Image)
		bg := stack.Objects[1].(*canvas.Rectangle)

		switch {
		case wallpapers[id].IsSFW():
			bg.StrokeColor = green
		case wallpapers[id].IsSketchy():
			bg.StrokeColor = yellow
		case wallpapers[id].IsNSFW():
			bg.StrokeColor = red
		}
		bg.Refresh()

		setImg := func(resource fyne.Resource) {
			fyne.Do(func() {
				img.Resource = resource
				img.Refresh()
			})
		}

		if value, ok := thumbCache[url]; ok {
			setImg(value)
			return
		}

		setImg(nil)

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

			setImg(res)
			fyne.Do(func() {
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
