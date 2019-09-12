package main

import (
	"github.com/fogleman/gg"
)

func main() {
	ggContext := gg.NewContext(1024, 1024)
	ggContext.DrawCircle(500, 500, 400)
	ggContext.SetRGB(0, 0, 0)
	ggContext.Fill()
	ggContext.SavePNG("circle.png")

	ggContext = gg.NewContext(1024, 1024)
	ggContext.SetRGB(1, 1, 1)
	ggContext.Clear()
	ggContext.SetRGB(0, 0, 0)
	if err := ggContext.LoadFontFace("/Library/Fonts/Arial.ttf", 100); err != nil {
		panic(err)
	}
	ggContext.DrawStringAnchored("Hello, world!", 500, 500, 0.5, 0.5)
	ggContext.SavePNG("text.png")
}
