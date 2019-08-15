package main

import (
	"fmt"

	"github.com/derekparker/trie"
)

func main() {
	t := trie.New()

	t.Add("helloworld", 1)
	t.Add("worldhello", 2)

	node, ok := t.Find("helloworld")
	meta := node.Meta()
	fmt.Println(meta, ok)

	helloResult1 := t.PrefixSearch("hello")
	fmt.Println(helloResult1)
	worldResult1 := t.PrefixSearch("world")
	fmt.Println(worldResult1)

	wrongResult1 := t.PrefixSearch("keyword")
	fmt.Println(wrongResult1)

	helloResult2 := t.HasKeysWithPrefix("hello")
	fmt.Println(helloResult2)
	worldResult2 := t.HasKeysWithPrefix("world")
	fmt.Println(worldResult2)

	wrongResult2 := t.HasKeysWithPrefix("keyword")
	fmt.Println(wrongResult2)

	Fuzzyresult1 := t.FuzzySearch("hw")
	fmt.Println(Fuzzyresult1)

	Fuzzyresult2 := t.FuzzySearch("wh")
	fmt.Println(Fuzzyresult2)
}
