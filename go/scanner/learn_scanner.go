package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	var source = `
// package.
package main

// import
import (
	"fmt"
)

// main
func main() {
	fmt.println("hello, world!")
}`

	var scannerVar scanner.Scanner
	scannerVar.Init(strings.NewReader(source))
	scannerVar.Filename = "example.go"

	for tok := scannerVar.Scan(); tok != scanner.EOF; tok = scannerVar.Scan() {
		fmt.Printf("%s: %s\n", scannerVar.Position, scannerVar.TokenText())
	}
}
