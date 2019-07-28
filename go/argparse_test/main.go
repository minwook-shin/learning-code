package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("test_program", "test argparse program")

	s := parser.String("v", "value", &argparse.Options{Required: true, Help: "Value to print"})
	f := parser.Flag("b", "bool", &argparse.Options{Required: true, Help: "Bool to print"})
	l := parser.List("l", "list", &argparse.Options{Required: true, Help: "List to print"})
	S := parser.Selector("s", "seletor", []string{"PRODUCT", "TEST"}, &argparse.Options{Required: true, Help: "Select to print"})
	F := parser.File("f", "file", os.O_RDWR, 0600, &argparse.Options{Required: true, Help: "File to print"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	fmt.Println(*s, *f, *l, *S, *F)
}
