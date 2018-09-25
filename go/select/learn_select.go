package main

import (
	"time"
	"fmt"
)

func function(c chan int)  {
	time.Sleep(1)
	c <- 21
}

func test1(c chan int)  {
	c <- 21
}
func test2(c chan int)  {
	c <- 12
}

func main() {
	c :=make(chan int)
	go function(c)
	// for {
	// 	select {
	// 		case value := <-c:
	// 			fmt.Println("received value: ", value)
	// 			return
	// 		default:
	// 			fmt.Println("no value received")
	// 		}
	// 	}
	
	c1 := make(chan int)
	c2 := make(chan int)
	go test1(c1)
	go test2(c2)
	
	time.Sleep(1)
	select{
	case v1 := <- c1:
		fmt.Println(v1)
	case v2 := <- c2:
		fmt.Println(v2)
	}

}