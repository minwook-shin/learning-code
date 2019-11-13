package main

import (
	"fmt"

	"github.com/wesovilabs/koazee"
)

func main() {
	intArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intArray)

	stream := koazee.StreamOf(intArray)
	fmt.Println(stream.Out().Val())

	stream = koazee.StreamOf(intArray)
	fmt.Println(stream.At(1).Int())

	fmt.Println(stream.First().Int())

	fmt.Println(stream.Last().Int())

	stream = koazee.StreamOf(intArray)
	fmt.Println(stream.Add(10).Do().Out().Val())

	fmt.Println(stream.Drop(1).Do().Out().Val())

	fmt.Println(stream.DropWhile(
		func(element int) bool {
			return element <= 5
		}).Do().Out().Val())

	fmt.Println(stream.DeleteAt(0).Do().Out().Val())

	fmt.Println(stream.Set(0, 10).Do().Out().Val())

	output, returnStream := stream.Pop()
	fmt.Println(output.Int())
	fmt.Println(returnStream.Out().Val())

	stream = koazee.StreamOf(intArray)
	count, _ := stream.Count()
	fmt.Println(count)

	indexOf, _ := stream.IndexOf(1)
	fmt.Println(indexOf)

	indexesOf, _ := stream.IndexesOf(1)
	fmt.Println(indexesOf)

	lastIndexOf, _ := stream.LastIndexOf(1)
	fmt.Println(lastIndexOf)

	contains, _ := stream.Contains(1)
	fmt.Println(contains)

	stream = koazee.StreamOf(intArray)
	fmt.Println(stream.Reverse().Out().Val())

	stream = koazee.StreamOf(intArray)
	fmt.Println(stream.Take(1, 3).Out().Val())

	fmt.Println(stream.
		Filter(
			func(val int) bool {
				return val > 3
			}).
		Out().Val(),
	)

	fmt.Println(stream.RemoveDuplicates().Out().Val())

	stream = koazee.StreamOf(intArray)
	value, _ := stream.GroupBy(func(val int) int { return val * 10 })
	fmt.Println(value)

	stream = koazee.StreamOf(intArray)
	fmt.Println(stream.Reduce(func(acc, val int) int {
		return acc + val
	}).Int())

	stream = koazee.StreamOf(intArray)
	stream.ForEach(func(i int) {
		fmt.Println(i)
	}).Do()

	stream = koazee.
		StreamOf(intArray).
		Filter(func(i int) bool { return i > 3 }).
		ForEach(func(i int) {
			fmt.Println(i)
		})
	stream.Do()
}
