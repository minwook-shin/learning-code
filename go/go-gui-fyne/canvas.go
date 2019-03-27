package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func main() {
	app := app.New()
	w := app.NewWindow("Canvas")

	content := fyne.NewContainer(
		canvas.NewText("Resize me", color.RGBA{0, 0x80, 0, 0xff}),

		&canvas.Rectangle{FillColor: color.RGBA{0, 0, 0x80, 0xff},
			StrokeColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
			StrokeWidth: 2},

		&canvas.Circle{StrokeColor: color.RGBA{0, 0, 0x80, 0xff},
			FillColor:   color.RGBA{0x30, 0x30, 0x30, 0x60},
			StrokeWidth: 2},

		&canvas.Line{StrokeColor: color.RGBA{0, 0, 0x80, 0xff}, StrokeWidth: 5},
	)
	content.Layout = layout.NewFixedGridLayout(fyne.NewSize(70, 70))

	w.SetContent(content)
	w.ShowAndRun()
}
