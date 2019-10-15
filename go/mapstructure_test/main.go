package main

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

func main() {
	type Book struct {
		Title string
		Score int
	}

	type Author struct {
		Name    string
		Age     int
		Created time.Time
		Book
	}

	author := map[string]interface{}{
		"Name":    "author_name_1",
		"Book":    Book{"book_name_1", 80},
		"Age":     "22",
		"Created": time.Now().UTC().Local(),
	}

	var result Author
	config := &mapstructure.DecoderConfig{
		ErrorUnused:      true,
		ZeroFields:       true,
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(author)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
