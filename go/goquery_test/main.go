package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("https://minwook-shin.github.io")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		panic(response.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	document.Find(".posts article .entry").Each(func(i int, s *goquery.Selection) {
		titleText := s.Find("p").Text()
		fmt.Println(titleText)
	})
}
