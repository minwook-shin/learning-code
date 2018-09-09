package main

import (
	"fmt"
	_ "package/calc"
)

func init(){
	fmt.Println("main init")
}

func main() {
	fmt.Println("main")
	// fmt.Println(calc.Sum(1.0,1.0))
	// fmt.Println(calc.Sub(1.0,1.0))
	// fmt.Println(calc.Div(1.0,1.0))
	// fmt.Println(calc.Mul(1.0,1.0))
}