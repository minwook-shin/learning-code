package main

import (
	"fmt"

	"github.com/peterbourgon/diskv"
)

func main() {
	flatTransform := func(s string) []string { return []string{} }

	diskv := diskv.New(diskv.Options{
		BasePath:     "data-dir",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})

	diskv.Write("byteKey", []byte{'h', 'e', 'l', 'l', 'o'})

	value, err := diskv.Read("byteKey")
	if err != nil {
		panic(err)
	}

	diskv.WriteString("stringKey", "hello")
	fmt.Println(diskv.ReadString("stringKey"))

	fmt.Printf("%v\n", value)

	diskv.Erase("byteKey")
	diskv.Erase("stringKey")
}
