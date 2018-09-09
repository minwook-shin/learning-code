package main

import (
	"unsafe"
	"reflect"
	"fmt"
)

func main() {
	a := true
	b := false
	fmt.Println(a,b)
	fmt.Println(a&&b)
	fmt.Println(a||b)

	var c int64 
	c = 9223372036854775807
	d := 9223372036854775807
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))

	var e uint
	e = 18446744073709551615
	fmt.Println(unsafe.Sizeof(e))

	f,g := 3.14, 3.15
	fmt.Println(f,g)
	tmp1 := f+g
	tmp2 := f-g
	fmt.Println(tmp1,tmp2)

	tmp3 := complex(1,5)
	tmp4 := 1+5i
	fmt.Println(tmp3,tmp4)

	var z = 'A'
	fmt.Println(z)

	var name string 
	name = "string"
	fmt.Println(name)

	var(i ,j =1, 3.14)
	fmt.Println(i + int(j))
}