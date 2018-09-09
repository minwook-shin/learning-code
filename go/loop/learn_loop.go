package main

import (
	"fmt"
)

func main() {
	count := 10
	for index := 0; index < count; index++ {
		fmt.Print(index)
	}

	for index := 0; index < count; index++ {
		if index == 5 {
			break
		}
		fmt.Print(index)
	}

	for index := 0; index < count; index++ {
		if index % 3 == 0 {
			continue
		}
		fmt.Print(index)
	}

	for ; count < 10 ; {
		
	}
	for count < 10 {
		
	}

	for {
		fmt.Println("hello?")
	}
}