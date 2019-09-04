package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

func main() {
	file := jen.NewFile("main")
	file.Func().Id("main").Params().Block(
		jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
	)
	fmt.Printf("%#v", file)
}
