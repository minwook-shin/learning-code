package main

import (
	"fmt"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/segment"
	"github.com/anthonynsimon/bild/transform"
)

func main() {
	image, err := imgio.Open("original.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	rotate := transform.Rotate(image, 40, nil)

	if err := imgio.Save("rotate.png", rotate, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

	brightness := adjust.Brightness(image, 0.20)
	if err := imgio.Save("brightness.png", brightness, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

	contrast := adjust.Contrast(image, -0.5)
	if err := imgio.Save("contrast.png", contrast, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

	hue := adjust.Hue(image, -40)
	if err := imgio.Save("hue.png", hue, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
	box := blur.Box(image, 8.0)
	if err := imgio.Save("box.png", box, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

	edgeDetection := effect.EdgeDetection(image, 1.0)
	if err := imgio.Save("edgeDetection.png", edgeDetection, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

	emboss := effect.Emboss(image)
	if err := imgio.Save("emboss.png", emboss, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

	grayscale := effect.Grayscale(image)
	if err := imgio.Save("grayscale.png", grayscale, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
	invert := effect.Invert(image)
	if err := imgio.Save("invert.png", invert, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
	sepia := effect.Sepia(image)
	if err := imgio.Save("sepia.png", sepia, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
	sobel := effect.Sobel(image)
	if err := imgio.Save("sobel.png", sobel, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
	threshold := segment.Threshold(image, 120)
	if err := imgio.Save("threshold.png", threshold, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}
