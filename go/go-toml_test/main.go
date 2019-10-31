package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func main() {
	type TestData struct {
		Default string `toml:"default"`
		Comment string `toml:"comment" commented:"true" comment:"hello world"`
		Data    string `toml:"data" commented:"false" comment:"hello world"`
	}
	type TestSection struct {
		TestData `toml:"Section" comment:"Test Section"`
	}

	section := TestSection{TestData{Default: "default_id", Comment: "test comment", Data: "test data"}}

	byteData, err := toml.Marshal(section)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byteData))

	section = TestSection{}
	toml.Unmarshal(byteData, &section)
	fmt.Println(section)
}
