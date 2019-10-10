package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
)

func main() {
	backoff := 2 * time.Millisecond
	maxJitter := 5 * time.Millisecond
	retrier := heimdall.NewRetrier(heimdall.NewConstantBackoff(backoff, maxJitter))

	timeout := 100 * time.Millisecond

	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(3))

	response, err := client.Get("http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))
}
