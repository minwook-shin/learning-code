package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestFunc(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/test").
		Reply(200).
		JSON(map[string]string{"hello": "world"})

	res, _ := http.Get("http://example.com/test")
	fmt.Println(res.StatusCode)

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	fmt.Println(gock.IsDone())

	gock.New("http://example.com").
		MatchHeader("Authorization", "^server$").
		MatchHeader("API", "[0-9]+").
		HeaderPresent("Accept").
		Reply(200).
		BodyString("hello,world!")

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Authorization", "server")
	req.Header.Set("API", "1")
	req.Header.Set("Accept", "text/plain")

	res, _ = (&http.Client{}).Do(req)
	fmt.Println(res.StatusCode)
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	fmt.Println(gock.IsDone())
}
