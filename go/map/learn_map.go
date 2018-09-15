package main

import (
	"fmt"
)

func main() {
	// var a map[string]int
	
	var a = make(map[string]int)
	a["key1"] = 1
	a["key2"] = 2
	fmt.Println(len(a),a)


	var b = map[string]int{
		"key1" : 1,
		"key2" : 2,
	}
	fmt.Println(b)

	var db = map[string]int{
		"key1" : 1,
		"key2" : 2,
	}
	fmt.Println(db["key1"])
	fmt.Println(db["key3"])
	v1 , status := db["key1"]
	fmt.Println(v1, status)
	v2 , status := db["key3"]
	fmt.Println(v2, status)

	forMap := map[string]int{
		"key1" : 1,
		"key2" : 2,
	}
	for i,j := range forMap{
		fmt.Println(i,j)
	}

	var findMap = map[string]int{
		"key1" : 1,
		"key2" : 2,
	}
	fmt.Println(findMap["key1"])
	delete(findMap,"key1")
	fmt.Println(findMap["key1"])
}