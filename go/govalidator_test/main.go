package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type testStruct struct {
	Title string `valid:"-"`
	Email string `valid:"email"`
	URL   string `valid:"url"`
	IP    string `valid:"ip"`
}

func main() {
	str := govalidator.ToString(&testStruct{Email: "test@example.com", IP: "127.0.0.1"})
	fmt.Println(str)

	test := testStruct{Email: "test@example.com", IP: "127.0.0.1"}
	fmt.Println(test)

	result, err := govalidator.ValidateStruct(test)
	if err != nil {
		fmt.Println(err)
	}
	println(result)

	println(govalidator.IsURL(`https://example.com`))

	intArray := []interface{}{1, 2, 3}
	var function govalidator.Iterator = func(v interface{}, _ int) {
		fmt.Println(v)
	}
	govalidator.Each(intArray, function)

	govalidator.TagMap["test"] = govalidator.Validator(func(s string) bool {
		return s == "test"
	})

	fmt.Println(govalidator.WhiteList("h1e2l3l4o5w6o7r8l9d0", "a-z"))
}
