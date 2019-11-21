package main

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func main() {
	uuid1, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	fmt.Println(uuid1)

	gen := uuid.NewGen()
	uuid2, err := gen.NewV4()
	if err != nil {
		panic(err)
	}
	fmt.Println(uuid2)

	var uuid3 = uuid.Must(uuid.NewV4())
	if err != nil {
		panic(err)
	}
	fmt.Println(uuid3)

	uuidStr := "1c880bea-a9af-41b7-ba89-f0b08ae3b741"
	uuid4, err := uuid.FromString(uuidStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(uuid4)
}
