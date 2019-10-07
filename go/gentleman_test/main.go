package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/compression"
)

func main() {
	client := gentleman.New()

	client.Use(compression.Disable())
	client.URL("http://httpbin.org")

	request := client.Request()
	request.Path("/headers")
	request.SetHeader("Custom-Header", "Custom-VALUE")

	response, err := request.Send()
	if err != nil {
		panic(err)
	}

	if !response.Ok {
		fmt.Println(response.StatusCode)
		return
	}

	fmt.Println(response.String())
}
