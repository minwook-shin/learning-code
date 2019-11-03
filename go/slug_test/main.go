package main

import (
	"fmt"

	"github.com/gosimple/slug"
)

func main() {
	slugText := slug.Make("Hellö Wörld")
	fmt.Println(slugText)

	slugText = slug.Make("고")
	fmt.Println(slugText)

	slugText = slug.MakeLang("python & go", "en")
	fmt.Println(slugText)

	slug.Lowercase = false
	slug.MaxLength = 20
	slugText = slug.MakeLang("python & go", "en")
	fmt.Println(slugText)

	slug.CustomSub = map[string]string{
		"world": "hello",
	}
	slugText = slug.Make("hello world")
	fmt.Println(slugText)

	isSlug := slug.IsSlug("hello-world")
	fmt.Println(isSlug)
}
