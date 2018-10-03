package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("======")

	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("======")

	tmp1 := list.PushBack(5)
	list.InsertAfter(4,tmp1)
	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("======")

	tmp2 := list.PushBack(5)
	list.InsertBefore(6,tmp2)
	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
