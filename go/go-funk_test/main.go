package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

type test struct {
	ID       int
	Password string
}

func main() {
	contains := funk.Contains([]string{"hello", "world"}, "hello")
	fmt.Println(contains)

	testStruct := &test{
		ID:       1,
		Password: "0000",
	}
	contains = funk.Contains([]*test{testStruct}, testStruct)
	fmt.Println(contains)
	contains = funk.Contains([]*test{testStruct}, nil)
	fmt.Println(contains)

	testStruct2 := &test{
		ID:       2,
		Password: "1111",
	}

	get := funk.Get(testStruct2, "ID")
	fmt.Println(get)

	contains = funk.Contains([]*test{testStruct}, testStruct2)
	fmt.Println(contains)

	contains = funk.Contains("hello", "lo")
	fmt.Println(contains)
	contains = funk.Contains("hello", "world")
	fmt.Println(contains)

	contains = funk.Contains(map[int]string{1: "test"}, 1)
	fmt.Println(contains)

	indexOf := funk.IndexOf([]string{"hello", "world"}, "hello")
	fmt.Println(indexOf)
	indexOf = funk.IndexOf([]string{"hello", "world"}, "test")
	fmt.Println(indexOf)

	indexOf = funk.LastIndexOf([]string{"hello", "hello", "world"}, "hello")
	fmt.Println(indexOf)
	indexOf = funk.LastIndexOf([]string{"hello", "world"}, "test")
	fmt.Println(indexOf)

	results := []*test{testStruct, testStruct2}
	mapping := funk.ToMap(results, "ID")
	fmt.Println(mapping)

	filter := funk.Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(filter)

	find := funk.Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(find)

	intMap := funk.Map([]int{1, 2, 3, 4}, func(x int) int {
		return x * 2
	})
	fmt.Println(intMap)

	stringMap := funk.Map([]int{1, 2, 3, 4}, func(x int) string {
		return "Hello"
	})
	fmt.Println(stringMap)

	intMap = funk.Map([]int{1, 2, 3, 4}, func(x int) (int, int) {
		return x, x
	})
	fmt.Println(intMap)

	mapping = map[int]string{
		1: "hello",
		2: "world",
	}

	intMap = funk.Map(mapping, func(k int, v string) int {
		return k
	})
	fmt.Println(intMap)

	stringMap = funk.Map(mapping, func(k int, v string) (string, string) {
		return fmt.Sprintf("%d", k), v
	})
	fmt.Println(stringMap)

	funk.ForEach([]int{1, 2, 3, 4}, func(x int) {
		fmt.Println(x)
	})

	keys := funk.Keys(map[string]int{"hello": 1, "world": 2})
	fmt.Println(keys)
	keys = funk.Keys(testStruct)
	fmt.Println(keys)

	values := funk.Values(map[string]int{"hello": 1, "world": 2})
	fmt.Println(values)
	values = funk.Values(testStruct)
	fmt.Println(values)

	chunk := funk.Chunk([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(chunk)

	flattenDeep := funk.FlattenDeep([][]int{[]int{1, 2}, []int{3, 4}})
	fmt.Println(flattenDeep)

	uniq := funk.Uniq([]int{0, 1, 1, 2, 2, 3, 3, 10})
	fmt.Println(uniq)

	drop := funk.Drop([]int{0, 0, 0, 0}, 3)
	fmt.Println(drop)

	intial := funk.Initial([]int{0, 1, 2, 3, 4})
	fmt.Println(intial)

	tail := funk.Tail([]int{0, 1, 2, 3, 4})
	fmt.Println(tail)

	shuffle := funk.Shuffle([]int{0, 1, 2, 3, 4})
	fmt.Println(shuffle)

	sum := funk.Sum([]int{0, 1, 2, 3, 4})
	fmt.Println(sum)

	reverse := funk.Reverse([]int{0, 1, 2, 3, 4})
	fmt.Println(reverse)

	slice := funk.SliceOf(testStruct)
	fmt.Println(slice)

	random := funk.RandomInt(0, 100)
	fmt.Println(random)
	random2 := funk.RandomString(4)
	fmt.Println(random2)

	shard := funk.Shard("e89d66bdfdd4dd26b682cc77e23a86eb", 2, 2, false)
	fmt.Println(shard)
	shard = funk.Shard("e89d66bdfdd4dd26b682cc77e23a86eb", 2, 2, true)
	fmt.Println(shard)
}
