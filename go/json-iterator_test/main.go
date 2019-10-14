package main

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	type Book struct {
		Title string
		Score int
	}

	type Author struct {
		Name string
		Book
		Age     int `csv:"age,omitempty"`
		Created time.Time
	}

	author := []Author{
		{
			Name:    "author_name_1",
			Book:    Book{"book_name_1", 80},
			Age:     22,
			Created: time.Now().UTC().Local(),
		},
		{
			Name: "author_name_2",
			Book: Book{"book_name_2", 70},
		},
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	byteData, err := json.Marshal(&author)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byteData))

	var authorCopy []Author
	err = json.Unmarshal(byteData, &authorCopy)
	if err != nil {
		panic(err)
	}

	for _, authorData := range authorCopy {
		fmt.Println(authorData)
	}
}
