package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	example := "hello, world!"
	if example == "hello, world!" {
		fmt.Println("hello")
	} else{
		fmt.Println("hello")
	}

	if temp := 3.14; temp > 3 {
		fmt.Println("true!")
	}else{
		fmt.Println("false")
	}

	if temp2 := 3.14; temp2 < 3 {
		fmt.Println("true!")
	}else if temp2 == 3.14{
		fmt.Println("pi!")
	}else{
		fmt.Println("false")
	}

	if text, err := ioutil.ReadFile("./test.txt");err != nil {
		fmt.Println(text,err)
	}
}