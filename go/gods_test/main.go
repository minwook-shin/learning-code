package main

import (
	"fmt"

	"github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/emirpasic/gods/trees/avltree"
	"github.com/emirpasic/gods/utils"
)

func main() {
	list := singlylinkedlist.New()
	list.Add("b", "a")
	list.Sort(utils.StringComparator)
	fmt.Println(list.Get(0))
	result := list.Contains("a", "b")
	fmt.Println(result)
	list.Swap(0, 1)
	list.Clear()
	fmt.Println(list.Empty(), list.Size())
	list.Insert(0, "b")
	list.Insert(0, "a")
	it := list.Iterator()
	for it.Next() {
		index, value := it.Index(), it.Value()
		fmt.Println(index, value)
	}

	set := hashset.New()
	set.Add(1, 2, 2, 3, 4, 5)
	set.Remove(4)
	fmt.Println(set.Values())
	set.Contains(1, 5)
	set.Clear()
	fmt.Println(set.Empty(), set.Size())

	stack := linkedliststack.New()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Values())
	_, _ = stack.Peek()
	_, _ = stack.Pop()
	fmt.Println(stack.Values())
	stack.Clear()
	fmt.Println(stack.Empty(), stack.Size())

	hashMap := hashmap.New()
	hashMap.Put(1, "x")
	hashMap.Put(2, "b")
	hashMap.Put(1, "a")
	fmt.Println(hashMap.Values(), hashMap.Keys())
	hashMap.Remove(1)
	hashMap.Clear()
	fmt.Println(hashMap.Empty(), hashMap.Size())

	tree := avltree.NewWithIntComparator()
	tree.Put(1, "a")
	tree.Put(2, "b")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(5, "e")
	fmt.Println(tree)
	fmt.Println(tree.Values())
	it2 := tree.Iterator()
	for it2.Next() {
		key, value := it2.Key(), it2.Value()
		fmt.Println(key, value)
	}
	tree.Clear()
	fmt.Println(tree.Empty(), tree.Size())
}
