package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

func main() {
	f := jen.NewFile("main")
	f.Func().Id("main").Params().Block(
		jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
	)
	fmt.Printf("%#v", f)
}
