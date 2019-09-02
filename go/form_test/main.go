package main

import (
	"fmt"
	"net/url"

	"github.com/go-playground/form"
)

type Test struct {
	Id           string `form:"id"`
	Age          uint8  `form:"age"`
	Email        string `form:"email"`
	Active       bool   `form:"active"`
	MapExample   map[string]string
	ArrayExample []string
}

var decoder *form.Decoder

var encoder *form.Encoder

func getValues() url.Values {
	return url.Values{
		"id":              []string{"test_id"},
		"age":             []string{"20"},
		"email":           []string{"test@example.com"},
		"active":          []string{"true"},
		"MapExample[key]": []string{"value"},
		"ArrayExample[0]": []string{"value"},
	}
}

func main() {
	decoder = form.NewDecoder()

	urlValues := getValues()

	var entity Test

	err := decoder.Decode(&entity, urlValues)
	if err != nil {
		panic(err)
	}

	fmt.Println(entity)

	encoder = form.NewEncoder()

	// entity = Test{
	// 	Id:           "test_id",
	// 	Age:          20,
	// 	Email:        "Male",
	// 	Active:       true,
	// 	MapExample:   map[string]string{"key": "value"},
	// 	ArrayExample: []string{"value"},
	// }

	urlValues, err = encoder.Encode(&entity)
	if err != nil {
		panic(err)
	}

	fmt.Println(urlValues)
}
