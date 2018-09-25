package main

import (  
	"time"
    "fmt"
	"sync"
)

func function(){
	
}

func main() {
	var m sync.Mutex
	value := 0

	go func() {
		for i := 0; i < 3; i++ {
			m.Lock()
			value = value + 10 
			fmt.Println(value)
			m.Unlock()
		}
	}()

	go func() {
		for j := 0; j < 2; j++ {
			m.Lock()
			value = value - 10
			fmt.Println(value)
			m.Unlock()
		}
	}()
	time.Sleep(1000000000)
}