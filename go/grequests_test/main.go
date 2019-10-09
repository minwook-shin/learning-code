package main

import (
	"fmt"

	"github.com/levigross/grequests"
)

func main() {
	response, err := grequests.Get("http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())

	requestOptions := &grequests.RequestOptions{
		Params: map[string]string{"q": "value"},
	}
	response, err = grequests.Get("http://httpbin.org/get?q=test", requestOptions)

	fmt.Println(response.String())

	if err := response.DownloadToFile("test.txt"); err != nil {
		panic(err)
	}

	response.ClearInternalBuffer()
	fmt.Println(response.String())
}
