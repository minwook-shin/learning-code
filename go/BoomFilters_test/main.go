package main

import (
	"fmt"

	boom "github.com/tylertreat/BoomFilters"
)

func main() {
	stableBloomFilter := boom.NewDefaultStableBloomFilter(10000, 0.01)
	point := stableBloomFilter.StablePoint()
	fmt.Println(point)

	stableBloomFilter.Add([]byte(`a`))

	if stableBloomFilter.Test([]byte(`a`)) {
		fmt.Println("A")
	}

	if stableBloomFilter.TestAndAdd([]byte(`b`)) {
		fmt.Println("B")
	} else {
		fmt.Println("not B")
	}

	if stableBloomFilter.Test([]byte(`b`)) {
		fmt.Println("B")
	}

	stableBloomFilter.Reset()

	first := []string{"hello", "world", "c++", "java", "python", "rust", "golang"}
	second := []string{"hello", "world", "c++", "java", "python"}

	similarity := boom.MinHash(first, second)
	fmt.Println(similarity)
}
