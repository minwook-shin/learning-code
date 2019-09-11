package main

import (
	"fmt"

	"github.com/pravj/geopattern"
)

func main() {
	args := map[string]string{}
	patternFile := geopattern.Generate(args)
	fmt.Println(patternFile)

	args = map[string]string{"color": "#f9b"}
	patternFile = geopattern.Generate(args)
	fmt.Println(patternFile)

	args = map[string]string{"baseColor": "#e2b"}
	patternFile = geopattern.Generate(args)
	fmt.Println(patternFile)

	args = map[string]string{"generator": "diamonds"}
	patternFile = geopattern.Generate(args)
	fmt.Println(patternFile)

	args = map[string]string{"phrase": "O"}
	patternFile = geopattern.Generate(args)
	fmt.Println(patternFile)

	args = map[string]string{}
	patternFile = geopattern.Base64String(args)
	fmt.Println(patternFile)

	args = map[string]string{}
	patternFile = geopattern.URIimage(args)
	fmt.Println(patternFile)
}
