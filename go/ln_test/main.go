package main

import "github.com/fogleman/ln/ln"

func main() {
	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	width := 1024.0
	height := 1024.0
	fovy := 50.0
	znear := 0.1
	zfar := 10.0
	step := 0.01

	paths := scene.Render(ln.Vector{3, 3, 3}, ln.Vector{0, 0, 0}, ln.Vector{0, 0, 1},
		width, height, fovy, znear, zfar, step)

	paths.WriteToPNG("WriteToPNG.png", width, height)
	paths.WriteToSVG("WriteToSVG.svg", width, height)
	paths.WriteToTXT("WriteToTXT.txt")

}
