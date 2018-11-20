package main

import(
	"fmt"
	"strings"
)

func main() {
	str := "hello, golang!"

	fmt.Println(strings.Contains(str,"go"))
	fmt.Println(strings.ContainsRune(str,'G'))

	fmt.Println(strings.Count(str,"l"))

	fmt.Println(strings.HasPrefix(str,"hello"))
	fmt.Println(strings.HasPrefix(str,"golang!"))

	fmt.Println(strings.HasSuffix(str,"hello"))
	fmt.Println(strings.HasSuffix(str,"golang!"))

	fmt.Println(strings.Index(str,"go"))
	fmt.Println(strings.IndexByte(str,'g'))

	fmt.Println(strings.Join([]string{"one","two"},"|"));

	fmt.Println(strings.Repeat(str, 3))

	fmt.Println(strings.Replace(str,"o","0",1))
	fmt.Println(strings.Replace(str,"o","0",-1))

	fmt.Println(strings.Split("go,c++,c,java,python",","))

	fmt.Println(strings.ToLower("HELLO, GOLANG!"))
	fmt.Println(strings.ToUpper(str))
}