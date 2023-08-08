package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ikhsan892/goceng"
)

type FyneAdapter struct {
}

func NewFyneAdapter(app goceng.App) *FyneAdapter {
	return &FyneAdapter{}
}

func (f *FyneAdapter) Init() error {
	a := app.New()

	w := a.NewWindow("")

	image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	// image := canvas.NewImageFromFile(fileName)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()

	return nil
}
