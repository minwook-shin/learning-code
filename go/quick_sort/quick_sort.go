package main

import (
	"fmt"
)

func partition(arr []int) int {
	pivot := arr[0]
	start, end := 0, len(arr)-1

	for {
		for start < len(arr) && arr[start] <= pivot {
			start++
		}
		for arr[end] > pivot {
			end--
		}
		if start >= end {
			arr[0], arr[end] = arr[end], arr[0]
			return end
		}

		arr[start], arr[end] = arr[end], arr[start]
	}
}

func quick(arr []int) {
	if len(arr) > 1 {
		p := partition(arr)
		quick(arr[0:p])
		quick(arr[p+1:])
	}
}

func main() {
	arr := []int{3, 7, 4, 1, 7, 5, 4, 8}
	quick(arr)
	fmt.Println(arr)
}
