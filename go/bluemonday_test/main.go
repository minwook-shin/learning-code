package main

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

func main() {
	policy := bluemonday.UGCPolicy()

	htmlString := policy.Sanitize(
		`<a onclick="alert("hello")" href="http://www.google.com">Google</a>`,
	)

	fmt.Println(htmlString)

	policy = bluemonday.NewPolicy()

	policy.AllowStandardURLs()

	policy.AllowAttrs("href").OnElements("a")
	policy.AllowElements("p")

	htmlString = policy.Sanitize(
		`<p><a onclick="alert("hello")" href="http://www.google.com">Google</a><p/>`,
	)

	fmt.Println(htmlString)
}
