package main

import (
	"fmt"
	"io/ioutil"

	"github.com/matcornic/hermes/v2"
)

func main() {
	hermesGenerator := hermes.Hermes{
		Product: hermes.Product{
			Name:        "minwook_test",
			Link:        "https://example.com/",
			Logo:        "https://octodex.github.com/images/original.png",
			Copyright:   "Copyright © 2019 test. All rights reserved.",
			TroubleText: "TroubleText",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: "Consumer",
			Intros: []string{
				"Welcome to example! Intros text",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Instructions, click here:",
					Button: hermes.Button{
						Text: "Click Text",
						Link: "https://example.com/",
					},
				},
			},
			Outros: []string{
				"Good bye, Outros.",
			},
			Greeting:  "Hello, world!",
			Signature: "Signature text",
			Table: hermes.Table{
				Data: [][]hermes.Entry{
					{
						{Key: "language", Value: "Golang"},
						{Key: "Description", Value: "Open source programming language"},
						{Key: "ETC", Value: ""},
					},
				},
			},
			Dictionary: []hermes.Entry{
				{Key: "fisrt", Value: "first value"},
				{Key: "second", Value: "second value"},
			},
		},
	}
	emailText, err := hermesGenerator.GeneratePlainText(email)
	if err != nil {
		panic(err)
	}
	fmt.Println(emailText)

	emailBody, err := hermesGenerator.GenerateHTML(email)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("index.html", []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}
}
