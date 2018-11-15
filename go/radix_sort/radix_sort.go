package main

import (
	"fmt"
)

func findLNum(arr []int) int {
	n := 0
	for _, j := range arr {
		if j > n {
			n = j
		}
	}
	return n
}

func radix(arr []int) []int {
	largestNum := findLNum(arr)
	digit := 1
	temp := make([]int, len(arr))

	for largestNum > digit {
		bucket := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for _, j := range arr {
			bucket[(j/digit)%10]++
		}
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := len(arr) - 1; i >= 0; i-- {
			bucket[(arr[i]/digit)%10]--
			temp[bucket[(arr[i]/digit)%10]] = arr[i]
		}
		copy(arr, temp)
		digit *= 10
	}
	return arr
}

func main() {
	arr := []int{9999, 5045, 5043, 145, 424, 613, 785, 835, 642, 413, 336, 253, 3000}
	fmt.Println(arr)
	fmt.Println(radix(arr))
}
