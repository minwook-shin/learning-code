package main

import (
	"fmt"
)

func bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{3, 7, 4, 1, 7, 5, 4, 8}
	fmt.Println(arr)
	fmt.Println(bubble(arr))
}
