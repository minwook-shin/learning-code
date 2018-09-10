package main

import (
	"fmt"
)

func main() {
	temp := 1
	switch temp {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	
	switch temp := 2 ; temp {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default :
		fmt.Println("hello,world!")
	}

	switch temp := 10 ; temp {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default :
		fmt.Println("hello,world!")
	}

	switch temp := 10 ; temp {
	case 0,10,20,30:
		fmt.Println("10?20?30?")
	default :
		fmt.Println("hello,world!")
	}

	switch temp := 3.14; {
	case temp > 0 && temp < 10:
		fmt.Println("0~10")
	case temp == 3.14:
		fmt.Println("3.14")
	}

	switch temp := 3.14; {
	case temp > 0 && temp < 10:
		fmt.Println("0~10")
		fallthrough
	case temp == 3.14:
		fmt.Println("3.14")
	}
}