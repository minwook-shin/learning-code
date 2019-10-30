package main

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/dustin/go-humanize/english"
)

func main() {

	mb := humanize.Bytes(5012345)
	fmt.Println(mb)

	t := time.Date(2018, time.November, 1, 0, 0, 0, 0, time.UTC)
	time := humanize.Time(t)
	fmt.Println(time)

	ordinal := humanize.Ordinal(99)
	fmt.Println(ordinal)

	comma := humanize.Comma(123456789)
	fmt.Println(comma)

	fmt.Println(humanize.Ftoa(2.00))

	fmt.Println(english.PluralWord(1, "keyboard", ""))
	fmt.Println(english.PluralWord(2, "keyboard", ""))

	fmt.Println(english.Plural(1, "keyboard", ""))
	fmt.Println(english.Plural(2, "keyboard", ""))

	fmt.Println(english.WordSeries([]string{"hello", "world"}, "and"))
}
