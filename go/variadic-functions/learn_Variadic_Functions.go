package main

import (
	"reflect"
	"fmt"
)

func testInt(n ... int)  {
	fmt.Println(reflect.TypeOf(n),n)
}

func testArray(n ... []int)  {
	fmt.Println(reflect.TypeOf(n),n)
}

func changeInt1(n ... []int){
	n[0][0] = 1
}

func changeInt2(n ... int){
	n[0] = 1
}

func main() {
	testInt(1,2,3)
	testInt(1,2,3,4,5)

	var a = make([]int,5,10)
	testArray(a)
	var b = make([]int,10,10)
	testArray(b)
	
	c := make([]int,3,3)
	changeInt1(c)
	fmt.Println(c)

	d := make([]int,3,3)
	changeInt2(d...)
	fmt.Println(d)
}