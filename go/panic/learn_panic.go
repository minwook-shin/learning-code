package main

import (
	"fmt"
)

func panicFunc()  {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("OPEN ERROR", r)
		}
	}()
	
	arr := []int{1,2,3}
	fmt.Println(arr[3])
}

func main() {
	// panic("panic!")
	
	panicFunc()
	fmt.Println("recover!")
}