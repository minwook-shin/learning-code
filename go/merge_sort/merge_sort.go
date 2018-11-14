package main

import (
	"fmt"
)

func sort(arr []int) []int {
	return divide(arr)
}

func divide(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	return merge(divide(arr[:middle]), divide(arr[middle:]))
}

func merge(l []int, r []int) []int {
	arr := make([]int, len(l)+len(r))

	for i, j, all := 0, 0, 0; all < len(l)+len(r); all++ {
		if i < len(l) && j >= len(r) {
			arr[all] = l[i]
			i++
			continue
		}
		if j < len(r) && i >= len(l) {
			arr[all] = r[j]
			j++
			continue
		}
		if l[i] < r[j] {
			arr[all] = l[i]
			i++
			continue
		}
		if l[i] >= r[j] {
			arr[all] = r[j]
			j++
			continue
		}
	}
	return arr
}

func main() {
	arr := []int{1, 4, 6, 7, 8, 6, 4, 3, 2}
	fmt.Println(arr)
	fmt.Println(sort(arr))
}
