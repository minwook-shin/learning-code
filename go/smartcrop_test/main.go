package main

import (
	"image"
	"image/png"
	"os"

	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
)

func main() {
	file, _ := os.Open("original.png")
	decodeImage, _, _ := image.Decode(file)

	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	crop, _ := analyzer.FindBestCrop(decodeImage, 250, 250)

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	subImage := decodeImage.(SubImager).SubImage(crop)
	png.Encode(file, subImage)
}
