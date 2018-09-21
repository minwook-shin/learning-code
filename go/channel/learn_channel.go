package main

import (
	"fmt"
)

func ch(c chan int)  {
	c <- 1
}

func ch2(c chan <- int)  {
	c <- 1
}

func ch3(c2 <- chan int)  {
	i:=<-c2
	fmt.Println(i)
}

func ch4(c chan int) {  
    for i := 0; i < 5; i++ {
        c <- i
    }
    close(c)
}

func main() {
	var a chan int
	fmt.Println(a)
	a = make(chan int)
	fmt.Println(a)

	c := make(chan int)
	go ch(c)
	i := <- c
	fmt.Println(i)
	
	c2 := make(chan int)
	go ch(c2)
	fmt.Println(<- c2)
	
	c3 := make(chan int)
	go ch(c3)
	go ch3(c3)

	c4 := make(chan int)
    go ch4(c4)
    for {
        v, ok := <-c4
        if ok == false {
            break
        }
        fmt.Println(v, ok)
    }
}