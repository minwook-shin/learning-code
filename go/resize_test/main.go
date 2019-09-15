package main

import (
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("original.png")
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	var width uint = 128
	var height uint = 128

	file, err = os.Create("NearestNeighbor.png")
	if err != nil {
		panic(err)
	}
	png.Encode(file, resize.Resize(width, height, img, resize.NearestNeighbor))

	file, err = os.Create("Bilinear.png")
	if err != nil {
		panic(err)
	}
	png.Encode(file, resize.Resize(width, height, img, resize.Bilinear))

	file, err = os.Create("Bicubic.png")
	if err != nil {
		panic(err)
	}
	png.Encode(file, resize.Resize(width, height, img, resize.Bicubic))

	file, err = os.Create("MitchellNetravali.png")
	if err != nil {
		panic(err)
	}
	png.Encode(file, resize.Resize(width, height, img, resize.MitchellNetravali))

	file, err = os.Create("Lanczos2.png")
	if err != nil {
		panic(err)
	}
	png.Encode(file, resize.Resize(width, height, img, resize.Lanczos2))

	file, err = os.Create("Lanczos3.png")
	if err != nil {
		panic(err)
	}
	png.Encode(file, resize.Resize(width, height, img, resize.Lanczos3))

	defer file.Close()
}
