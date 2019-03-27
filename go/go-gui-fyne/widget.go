package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/app"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func makeProgress() fyne.Widget {
	progress := widget.NewProgressBar()
	infProgress := widget.NewProgressBarInfinite()

	go func() {
		num := 0.0
		for num < 1.0 {
			time.Sleep(100 * time.Millisecond)
			progress.SetValue(num)
			num += 0.01
		}

		progress.SetValue(1)
	}()

	return widget.NewVBox(
		widget.NewLabel("Percent"), progress,
		widget.NewLabel("Infinite"), infProgress)
}

func main() {
	app := app.New()
	w := app.NewWindow("위젯")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entry")

	w.SetContent(widget.NewTabContainer(
		widget.NewTabItem("Buttons", widget.NewVBox(
			widget.NewButton("Text button", func() { fmt.Println("tapped") }),
			widget.NewButtonWithIcon("With icon", theme.ConfirmIcon(), func() { fmt.Println("tapped") }),
		)),
		widget.NewTabItem("Input", widget.NewVBox(
			entry,
			widget.NewCheck("Check", func(on bool) { fmt.Println(on, entry.Text) }),
			widget.NewRadio([]string{"1", "2"}, func(s string) { fmt.Println(s) }),
		)),
		widget.NewTabItem("Progress", makeProgress()),
	))

	w.ShowAndRun()
}
