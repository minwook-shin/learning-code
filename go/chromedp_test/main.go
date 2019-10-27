package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	contextVar, cancelFunc := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancelFunc()

	contextVar, cancelFunc = context.WithTimeout(contextVar, 10*time.Second)
	defer cancelFunc()

	var strVar string

	err := chromedp.Run(contextVar,
		chromedp.Navigate(`https://golang.org/pkg/fmt/`),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Click(`#pkg-examples > div`, chromedp.NodeVisible),
		chromedp.Value(`#example_Println .play .input textarea`, &strVar),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(strVar)
}
