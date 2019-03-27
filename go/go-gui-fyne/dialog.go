package main

import (
	"errors"
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	win := app.NewWindow("Dialogs")

	win.SetContent(widget.NewVBox(
		widget.NewButton("Info", func() {
			dialog.ShowInformation("타이틀", "information msg...", win)
		}),
		widget.NewButton("Error", func() {
			err := errors.New("error msg")
			dialog.ShowError(err, win)
		}),
		widget.NewButton("Confirm", func() {
			cnf := dialog.NewConfirm("확인", "Are you enjoying this program?", (func(res bool) { fmt.Println(res) }), win)
			cnf.SetDismissText("no")
			cnf.SetConfirmText("Yes")
			cnf.Show()
		}),
	))
	win.ShowAndRun()
}
