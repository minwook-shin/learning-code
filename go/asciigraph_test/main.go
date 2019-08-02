package main

import (
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	graph := asciigraph.Plot(data, asciigraph.Width(100), asciigraph.Caption("increase"))
	fmt.Println(graph)

	var data2 []float64

	for i := 0; i < 100; i++ {
		data2 = append(data2, math.Sin(float64(i)*((math.Pi*1)/10)))
	}
	graph2 := asciigraph.Plot(data2, asciigraph.Height(10), asciigraph.Caption("sin"))
	fmt.Print(graph2)
}
