package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func layout(gui *gocui.Gui) error {
	maxX, maxY := gui.Size()
	if view, err := gui.SetView("hello", maxX/2-5, maxY/2, maxX/2+2, maxY/2+2); err != nil {
		fmt.Fprintln(view, "Hello!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer gui.Close()

	gui.SetManagerFunc(layout)

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}
