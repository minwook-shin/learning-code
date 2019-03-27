package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("test project")

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello golang!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
