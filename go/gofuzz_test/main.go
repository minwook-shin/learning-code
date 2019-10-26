package main

import (
	"fmt"

	fuzz "github.com/google/gofuzz"
)

func main() {
	Fuzzer := fuzz.New()
	var intValue int
	Fuzzer.Fuzz(&intValue)
	fmt.Printf("%#v \n", intValue)

	Fuzzer = fuzz.New().NilChance(0).NumElements(1, 1)
	var mapValue map[int]string
	Fuzzer.Fuzz(&mapValue)
	fmt.Printf("%#v \n", mapValue)

	Fuzzer = fuzz.New().NilChance(1)
	var structValue struct {
		A, B *string
	}
	Fuzzer.Fuzz(&structValue)
	fmt.Printf("%#v \n", structValue)

	type testStruct struct {
		A, B *string
	}

	Fuzzer = fuzz.New().NilChance(0).Funcs(
		func(e *testStruct, c fuzz.Continue) {
			switch c.Intn(2) {
			case 0:
				c.Fuzz(&e.A)
			case 1:
				c.Fuzz(&e.B)
			}
		},
	)

	var testObject testStruct
	Fuzzer.Fuzz(&testObject)
	fmt.Printf("%#v \n", testObject)
}
