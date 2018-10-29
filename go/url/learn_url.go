package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	value := url.Values{}
	value.Set("name", "go")
	value.Add("friend", "c")
	value.Add("friend", "java")
	value.Add("friend", "python")
	fmt.Println(value.Encode())

	parseValue, _ := url.ParseQuery(`name=go&friend=c++&friend=java;ect`)
	jsonByte, _ := json.Marshal(parseValue)
	fmt.Println(string(jsonByte))

	url, _ := url.Parse("http://golang.org")
	url.Scheme = "https"
	url.Host = "google.com"
	queryValue := url.Query()
	queryValue.Set("q", "golang")
	url.RawQuery = queryValue.Encode()
	fmt.Println(url)
}
