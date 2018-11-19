package main

import (
	"fmt"
	"strings"
)

func indexFunc(str []string, t string) int {
	for i, j := range str {
		if j == t {
			return i
		}
	}
	return -1
}

func includeFunc(str []string, t string) bool {
	return indexFunc(str, t) >= 0
}

func anyFunc(str []string, f func(string) bool) bool {
	for _, j := range str {
		if f(j) {
			return true
		}
	}
	return false
}

func allFunc(str []string, f func(string) bool) bool {
	for _, j := range str {
		if !f(j) {
			return false
		}
	}
	return true
}

func filterFunc(str []string, f func(string) bool) []string {
	newArr := []string{}
	for _, j := range str {
		if f(j) {
			newArr = append(newArr, j)
		}
	}
	return newArr
}

func mapFunc(str []string, f func(string) string) []string {
	newArr := make([]string, len(str))
	for i, j := range str {
		newArr[i] = f(j)
	}
	return newArr
}

func main() {
	var strs = []string{"google", "apple", "microsoft", "ibm"}

	fmt.Println(indexFunc(strs, "google"))

	fmt.Println(includeFunc(strs, "google"))

	fmt.Println(anyFunc(strs, func(v string) bool {
		return strings.HasPrefix(v, "g")
	}))

	fmt.Println(allFunc(strs, func(v string) bool {
		return strings.HasPrefix(v, "g")
	}))

	fmt.Println(filterFunc(strs, func(v string) bool {
		return strings.Contains(v, "g")
	}))

	fmt.Println(mapFunc(strs, strings.ToUpper))
}
