package main

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ssh struct {
	Host     string
	User     string
	Port     string
	Password string
}

func (s ssh) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Host, validation.Required),
		validation.Field(&s.Port),
		validation.Field(&s.User, validation.Required),
		validation.Field(&s.Password),
	)
}

func main() {
	value := "not_url"
	err := validation.Validate(value, validation.Required, is.URL)
	fmt.Println(err)

	value = "not_digits"
	err = validation.Validate(value, validation.Match(regexp.MustCompile("^[0-9]$")).Error("must be a string with digits"))
	fmt.Println(err)

	s := ssh{
		Host: "127.0.0.1",
		Port: "8080",
	}
	err = s.Validate()
	fmt.Println(err)

	sliceValue := []ssh{
		ssh{Host: "127.0.0.1", Port: "8080"},
		ssh{Host: "127.0.0.1", User: "test", Password: "1234"},
		ssh{Host: "Unknown"},
	}
	err = validation.Validate(sliceValue)
	fmt.Println(err)
}
