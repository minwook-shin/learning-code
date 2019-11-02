package main

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	parser := gofeed.NewParser()
	feed, _ := parser.ParseURL("http://minwook-shin.github.io/feed")
	fmt.Println(feed)

	fmt.Println(feed.Title)
	fmt.Println(feed.Description)
	fmt.Println(feed.Items)

	feedData := `<rss version="2.0">
<channel>
<title>test</title>
</channel>
</rss>`
	parser = gofeed.NewParser()
	feed, _ = parser.ParseString(feedData)
	fmt.Println(feed.Title)

	file, _ := os.Open("test.xml")
	defer file.Close()
	parser = gofeed.NewParser()
	feed, _ = parser.Parse(file)
	fmt.Println(feed)

	fmt.Println(feed.Title)
	fmt.Println(feed.Description)
	fmt.Println(feed.Items)
}
