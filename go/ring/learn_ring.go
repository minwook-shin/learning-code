package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(3)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	fmt.Println("======")
	s := ring.New(3)
	for j := 0; j < s.Len(); j++ {
		s.Value = 4
		s = s.Next()
	}
	rs := r.Link(s)
	rs.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	fmt.Println("======")
	r = r.Move(3)
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
