package main

import (
	"image/png"
	"os"

	"github.com/fogleman/pt/pt"
)

func main() {
	scene := pt.Scene{}

	material := pt.DiffuseMaterial(pt.White)

	plane := pt.NewPlane(pt.V(0, 0, 0), pt.V(0, 0, 1), material)
	scene.Add(plane)

	sphere := pt.NewSphere(pt.V(0, 0, 1), 1, material)
	scene.Add(sphere)

	light := pt.NewSphere(pt.V(0, 0, 5), 1, pt.LightMaterial(pt.White, 10))
	scene.Add(light)

	camera := pt.LookAt(pt.V(3, 3, 3), pt.V(0, 0, 0.5), pt.V(0, 0, 1), 50)

	defaultSampler := pt.NewSampler(256, 256)
	renderer := pt.NewRenderer(&scene, &camera, defaultSampler, 800, 600)

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(file, renderer.Render())
	if err != nil {
		panic(err)
	}
}
