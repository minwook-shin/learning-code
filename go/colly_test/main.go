package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	collector := colly.NewCollector(
		colly.AllowedDomains("minwook-shin.github.io"),
		colly.MaxDepth(1),
		colly.Async(true),
	)

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(e.Text, link)
		collector.Visit(e.Request.AbsoluteURL(link))
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	collector.Visit("https://minwook-shin.github.io")
}
