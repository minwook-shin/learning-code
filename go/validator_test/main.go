package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type Account struct {
	ID       string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required,email"`
	Age      uint8  `validate:"gte=0,lte=100"`
}

func main() {
	validate := validator.New()

	account := &Account{
		ID:    "test_id",
		Email: "test@example.com",
		Age:   20,
	}

	err := validate.Struct(account)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return
	}

	email := "test@example.com"
	err = validate.Var(email, "required,email")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return
	}
}
