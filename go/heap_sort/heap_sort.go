package main

import "fmt"

func sort(arr []int) []int {
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	for i := len(arr); i > 1; i-- {
		arr[0], arr[i-1] = arr[i-1], arr[0]
		heapify(arr, 0, i-1)
	}
	return arr
}

func heapify(arr []int, root, length int) {
	var max = root
	var l, r = (root * 2) + 1, (root * 2) + 2

	if l < length && arr[l] > arr[max] {
		max = l
	}

	if r < length && arr[r] > arr[max] {
		max = r
	}

	if max != root {
		arr[root], arr[max] = arr[max], arr[root]
		heapify(arr, max, length)
	}
}

func main() {
	arr := []int{9999, 5045, 5043, 145, 424, 613, 785, 835, 642, 413, 336, 253, 3000}

	fmt.Println(arr)
	fmt.Println(sort(arr))
}
