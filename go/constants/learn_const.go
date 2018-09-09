package main

import (
	"math"
	"fmt"
)

func main() {
	const a = 10
	fmt.Println(a)

	const pi = math.Pi
	fmt.Println(pi)
	// const sqrt = math.Sqrt(1)

	const hello = "Hello"
	fmt.Println(hello)
	
	type str string
	var world str = "world"
	// world = string("")
	fmt.Println(world)

	const boolean = true
	fmt.Println(boolean)

	const c = 3
    v := c
    var v32 int32 = c
    var v64 float64 = c
    var c64 complex64 = c
    fmt.Println(v, v32, v64, c64)
}