package main

import "fmt"
import "sort"

type blog struct {
	Author string
	Text   string
	year   int
}

type customInterface []blog

func (c customInterface) Len() int {
	return len(c)
}
func (c customInterface) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c customInterface) Less(i, j int) bool {
	return c[i].year < c[j].year
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{3, 1, 2}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println(s)

	myBlog := []blog{
		{"example blog","test", 4},
		{"example blog","test", 3},
		{"example blog","test", 1},
		{"example blog","test", 2},
	}

	fmt.Println(myBlog)
	sort.Sort(customInterface(myBlog))
	fmt.Println(myBlog)
}
