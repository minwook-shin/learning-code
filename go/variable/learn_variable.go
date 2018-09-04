package main

import (
	"math"
	"fmt"
	
)

func main() {
	var count int
	fmt.Println(count)
	count = 1
	fmt.Println(count)

	// var tmp int = 10
	// fmt.Println(tmp)
	var tmp2 = 10
	fmt.Println(tmp2)

	var age,time int = 21, 1000
	fmt.Println(age,time)

	var age2,time2 = 21, 1000
	fmt.Println(age2,time2)

	var(c1,c2 int = 10,10)
	fmt.Println(c1,c2)

	name := "name"
	fmt.Println(name)

	name,hello := "world" , "hello"
	fmt.Println(name, hello)

	fmt.Println(math.Pi	)
}
