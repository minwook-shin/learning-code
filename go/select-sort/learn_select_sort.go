package main

import (
	"fmt"
)

func selection(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		m := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[m] {
				m = j
			}
		}
		arr[m], arr[i] = arr[i], arr[m]
	}
}

func main() {
	arr := []int{7,5,3,9,6,4}
	fmt.Println(arr)
	selection(arr)
	fmt.Println(arr)
}
