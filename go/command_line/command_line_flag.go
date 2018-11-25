package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("string", "default", "string")
	numPtr := flag.Int("int", 404, "int")
	boolPtr := flag.Bool("bool", true, "bool")

	var strVar string
	var intVar int
	var boolVar bool
	flag.StringVar(&strVar, "strVar", "default", "string")
	flag.IntVar(&intVar, "intVar", 404, "int")
	flag.BoolVar(&boolVar, "boolVar", true, "bool")

	flag.Parse()

	fmt.Println(*wordPtr, *numPtr, *boolPtr)
	fmt.Println(strVar, intVar, boolVar)
	fmt.Println(flag.Args())
}
