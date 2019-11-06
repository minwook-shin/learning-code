package main

import (
	"fmt"

	"github.com/mvdan/xurls"
)

func main() {
	rxRelaxed := xurls.Relaxed()
	relaxed := rxRelaxed.FindString("hello, world! test.org!")
	fmt.Println(relaxed)

	relaxed = rxRelaxed.FindString("This is not URL")
	fmt.Println(relaxed)

	rxStrict := xurls.Strict()
	strict := rxStrict.FindAllString("See http://example.com/.", -1)
	fmt.Println(strict)

	strict = rxStrict.FindAllString("hello, world! test.org!", -1)
	fmt.Println(strict)

	regexp, _ := xurls.StrictMatchingScheme(`^[a-z]+`)
	fmt.Println(regexp.MatchString("http://hello.com"))
}
