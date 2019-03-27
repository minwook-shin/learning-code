package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	win := app.NewWindow("Advanced")

	win.SetContent(widget.NewVBox(
		widget.NewButton("Fullscreen", func() {
			win.SetFullScreen(!win.FullScreen())
		}),
	))
	win.ShowAndRun()
}
