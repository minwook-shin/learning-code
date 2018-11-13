package main

import (
	"fmt"
)

func insertion(arr []int, h int) {
	for i := h; i < len(arr); i++ {
		for j := i; j > h-1 && arr[j-h] > arr[j]; j = j - h {
			arr[j], arr[j-h] = arr[j-h], arr[j]
		}
	}
}

func shell(arr []int) []int {
	h := 1
	for h < len(arr)/3 {
		h = 3*h + 1
	}
	for h > 0 {
		insertion(arr, h)
		h = (h - 1) / 3
	}
	return arr
}

func main() {
	arr := []int{3, 7, 4, 1, 7, 5, 4, 8}
	fmt.Println(arr)
	fmt.Println(shell(arr))
}
