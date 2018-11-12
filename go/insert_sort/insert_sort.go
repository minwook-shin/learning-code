package main

import (
	"fmt"
)

func insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}

func main() {
	arr := []int{3, 7, 4, 1, 7, 5, 4, 8}
	fmt.Println(arr)
	fmt.Println(insertion(arr))
}
