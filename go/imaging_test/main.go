package main

import (
	"github.com/disintegration/imaging"
)

func main() {
	src, err := imaging.Open("original.png")
	if err != nil {
		panic(err)
	}

	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	blur := imaging.Blur(src, 5)
	err = imaging.Save(blur, "blur.jpg")
	if err != nil {
		panic(err)
	}

	sharpen := imaging.Sharpen(src, 1.5)
	err = imaging.Save(sharpen, "sharpen.jpg")
	if err != nil {
		panic(err)
	}
	adjustGamma := imaging.AdjustGamma(src, 1.25)
	err = imaging.Save(adjustGamma, "adjustGamma.jpg")
	if err != nil {
		panic(err)
	}
	adjustContrast := imaging.AdjustContrast(src, 20)
	err = imaging.Save(adjustContrast, "adjustContrast.jpg")
	if err != nil {
		panic(err)
	}
	adjustBrightness := imaging.AdjustBrightness(src, 20)
	err = imaging.Save(adjustBrightness, "adjustBrightness.jpg")
	if err != nil {
		panic(err)
	}
	adjustSaturation := imaging.AdjustSaturation(src, -30)
	err = imaging.Save(adjustSaturation, "adjustSaturation.jpg")
	if err != nil {
		panic(err)
	}

	grayscale := imaging.Grayscale(src)
	err = imaging.Save(grayscale, "grayscale.jpg")
	if err != nil {
		panic(err)
	}
	invert := imaging.Invert(src)
	err = imaging.Save(invert, "invert.jpg")
	if err != nil {
		panic(err)
	}

	convolve3x3 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, 0, 1,
			-1, 0, 1,
			-1, 0, 1,
		},
		nil,
	)
	err = imaging.Save(convolve3x3, "convolve3x3.jpg")
	if err != nil {
		panic(err)
	}

}
