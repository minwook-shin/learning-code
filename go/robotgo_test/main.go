package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.TypeStr("Hello World")
	robotgo.TypeString("Hello World")

	robotgo.KeyTap("enter")

	robotgo.WriteAll("WriteAll")
	text, err := robotgo.ReadAll()
	if err == nil {
		fmt.Println(text)
	}

	robotgo.ScrollMouse(1, "up")

	robotgo.MoveMouseSmooth(0, 0)
	robotgo.MouseClick("left", true)

	x, y := robotgo.GetMousePos()
	fmt.Println(x, y)

	color := robotgo.GetPixelColor(x, y)
	fmt.Println(color)

	bitmap := robotgo.CaptureScreen(1, 1, 3000, 6000)
	robotgo.SaveBitmap(bitmap, "SaveBitmap.png")
	defer robotgo.FreeBitmap(bitmap)

	event := robotgo.AddEvent("mleft")
	if event {
		fmt.Println("press")
	}

	terminalPid, err := robotgo.FindIds("Terminal")

	isExist, err := robotgo.PidExists(terminalPid[0])
	if err == nil && isExist {
		robotgo.Kill(terminalPid[0])
	}

	isAlert := robotgo.ShowAlert("제목", "메시지")
	if isAlert == 0 {
		fmt.Println("ok")
	}
}
