package main

import (
	"fmt"

	"gopkg.in/thedevsaddam/gojsonq.v2"
)

func main() {
	const text = `{"entity":[
		{"id":"first",
		"description":[
			{"name":"golang","score":90},
			{"name":"java","score":50}]}
			]}`

	jsonSTR := gojsonq.New().JSONString(text)

	id := jsonSTR.Find("entity.[0].id")
	fmt.Println(id)

	name := jsonSTR.Reset().Find("entity.[0].description.[0].name")
	fmt.Println(name)
}
