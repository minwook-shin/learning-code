package main

import (
	"fmt"
	"time"

	"github.com/jszwec/csvutil"
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

	byteData, err := csvutil.Marshal(author)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byteData))

	var authorCopy []Author
	if err := csvutil.Unmarshal(byteData, &authorCopy); err != nil {
		panic(err)
	}

	for _, authorData := range authorCopy {
		fmt.Println(authorData)
	}
}
