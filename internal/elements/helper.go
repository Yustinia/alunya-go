package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func formItem(label string, container fyne.CanvasObject) *widget.Form {
	return widget.NewForm(
		widget.NewFormItem(label, container),
	)
}
