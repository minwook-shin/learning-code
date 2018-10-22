package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fileSet := token.NewFileSet()

	source := 
	`
	package main

	import (
		"fmt"
	)
	
	func main() {
		fmt.Println("hello, world!")
	}
	`

	file, _ := parser.ParseFile(fileSet, "", source, parser.ImportsOnly)
	
	fmt.Println(file.Name)
	for _, spec := range file.Imports {
		fmt.Println(spec.Path.Value)
	}
}