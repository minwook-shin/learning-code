package main

import (
	"fmt"

	"github.com/huandu/xstrings"
)

func main() {
	fmt.Println(xstrings.Center("hello", 10, "123"))

	fmt.Println(xstrings.Count("hello", "aeiou"))

	fmt.Println(xstrings.Delete("hello", "aeiou"))

	fmt.Println(xstrings.ExpandTabs("h\tell\to\twor\tld", 4))

	fmt.Println(xstrings.FirstRuneToLower("HELLO"))
	fmt.Println(xstrings.FirstRuneToUpper("hello"))

	fmt.Println(xstrings.Insert("hello", "world", 1))

	fmt.Println(xstrings.LastPartition("hello", "l"))

	fmt.Println(xstrings.LeftJustify("hello", 10, "11"))
	fmt.Println(xstrings.RightJustify("hello", 10, "11"))

	fmt.Println(xstrings.Len("hello"))

	fmt.Println(xstrings.Partition("hello", "l"))

	fmt.Println(xstrings.Reverse("hello"))

	fmt.Println(xstrings.RuneWidth('가'))

	fmt.Println(xstrings.Shuffle("hello"))

	fmt.Println(xstrings.Slice("hello", 1, 3))

	fmt.Println(xstrings.Squeeze("hello", "l"))

	fmt.Println(xstrings.Successor("hello1"))

	fmt.Println(xstrings.SwapCase("hELLO"))

	fmt.Println(xstrings.ToCamelCase("hello_world"))
	fmt.Println(xstrings.ToKebabCase("hello_world"))
	fmt.Println(xstrings.ToSnakeCase("HelloWorld"))

	fmt.Println(xstrings.Translate("hello", "^l", "*"))

	fmt.Println(xstrings.Width("hello world"))

	fmt.Println(xstrings.WordCount("hello_world"))

	fmt.Println(xstrings.WordSplit("hello_world"))
}
