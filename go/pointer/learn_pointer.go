package main

import (
	"reflect"
	"fmt"
)

func main() {
	a := 10
	b :=  &a
	
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(b),b)

	var nilPtr *int
	fmt.Println(nilPtr)

	v := 10
	ptr := &v
	fmt.Println(*ptr)

	v2 := 10
	ptr2 := &v2
	*ptr2++
	fmt.Println(*ptr2)

	v3 := 10
	ptr3 := &v3
	*ptr3 = 20
	fmt.Println(*ptr3)

	arr := [3]int{1,2,3}
	arrPtr := &arr
	fmt.Println(arrPtr)
}