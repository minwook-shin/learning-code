package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants"
)

var sum int32

func helloFunc() {
	fmt.Println("Hello World!")
}

func intFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Println(n)
}

func main() {
	var waitGroup sync.WaitGroup

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		_ = ants.Submit(func() {
			helloFunc()
			waitGroup.Done()
		})
	}
	waitGroup.Wait()

	poolFunc, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		intFunc(i)
		waitGroup.Done()
	})

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		_ = poolFunc.Invoke(int32(i))
	}
	waitGroup.Wait()
	fmt.Println(sum)

	defer ants.Release()
	defer poolFunc.Release()
}
