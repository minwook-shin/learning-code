package main

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type testStruct struct {
	Title  string
	Entity entityInfo
}

type entityInfo struct {
	Name        string `toml:"name"`
	Password    string `toml:"password"`
	Description string `toml:"description"`
	TestTIme    time.Time
}

func main() {
	var test testStruct
	_, err := toml.DecodeFile("test.toml", &test)
	if err != nil {
		panic(err)
	}

	fmt.Println(test.Title)
	fmt.Println(test.Entity.Name)
	fmt.Println(test.Entity.Password)
	fmt.Println(test.Entity.Description)
	fmt.Println(test.Entity.TestTIme)
}
