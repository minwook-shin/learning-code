package main

import (
	"fmt"
	"io/ioutil"
	"net/mail"
	"strings"
)

func main() {
	adress, _ := mail.ParseAddress("example <example@example.com>")
	fmt.Println(adress.Name, adress.Address)

	var adressList = "example1 <example@example.com>, example2 <example@example.com>"
	emails, _ := mail.ParseAddressList(adressList)
	for _, i := range emails {
		fmt.Println(i.Name, i.Address)
	}


	msg := `Date: 2018-10-28
From: example1 <from@example.com>
To: example2 <to@example.com>
Subject: test

Message body text
`

	reader := strings.NewReader(msg)
	message, _ := mail.ReadMessage(reader)
	header := message.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	bodyText, _ := ioutil.ReadAll(message.Body)
	fmt.Printf("%s", bodyText)
}
