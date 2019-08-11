package main

import (
	"fmt"
	"math/rand"

	"github.com/willf/bitset"
)

func main() {
	var bs bitset.BitSet

	num := uint(rand.Intn(50))
	bs.Set(num)
	fmt.Println(bs.String())

	bs.Flip(num)

	fmt.Print(bs.All())
	fmt.Print("\n")

	fmt.Print(bs.Any())
	fmt.Print("\n")

	fmt.Print(bs.None())
	fmt.Print("\n")

	fmt.Print(bs.Len())

	bs.ClearAll()
	fmt.Println(bs.String())

	fmt.Print(bs.All())
	fmt.Print("\n")

	fmt.Print(bs.Any())
	fmt.Print("\n")

	fmt.Print(bs.None())
	fmt.Print("\n")

	bs.Set(1).Set(2)
	fmt.Println(bs.String())

	for index, error := bs.NextSet(0); error; index, error = bs.NextSet(index + 1) {
		fmt.Println(index)
	}
	if bs.Intersection(bitset.New(100).Set(2)).Count() == 1 {
		fmt.Println("BitSet equivalent of '&'")
	}
}
