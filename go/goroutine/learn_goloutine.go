package main

import (
	"sync"
	"time"
	"fmt"
)

func helloworld(){
	fmt.Println("hello, world!")
}

func say1() {  
    for i := 0; i < 10; i++ {
		fmt.Println("hello!")
		time.Sleep(1 * time.Millisecond)
	}
}

func say2() {  
    for i := 0; i < 10; i++ {
		fmt.Println("go?")
		time.Sleep(1* time.Millisecond)
	}
}

func main() {
	go helloworld()
	time.Sleep(1 * time.Second)
	fmt.Println("hello?")

	go say1()
	go say2()
	time.Sleep(100 * time.Millisecond)


	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hello! go function!")
			time.Sleep(1 * time.Millisecond)
		}
	}()
	time.Sleep(10 * time.Millisecond)

	var w sync.WaitGroup
	w.Add(2)
	go func() {
		defer w.Done()
		for i := 0; i < 2; i++ {
		fmt.Println("hello!")
		}
	}()
	go func() {
		defer w.Done()
		for i := 0; i < 2; i++ {
		fmt.Println("go function!")
		}
	}()
	w.Wait()
}