package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.Create("test.txt")
	fmt.Fprintln(data,"파일에 쓰고 있습니다.")
	fmt.Fprintln(data,"Fprintln() 함수는 개행이 됩니다.")
	fmt.Fprintln(data,"Fprint(), Fprintf() 함수도 있습니다.")
	data.Close()

	var t string
	read, _ := os.Open("test.txt")
	for i := 0; i < 11; i++ {
		fmt.Fscan(read,&t)
		fmt.Print(t," ")
	}
	read.Close()
}