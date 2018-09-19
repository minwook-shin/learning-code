package main

import (
	"fmt"
)

type custom string

type sample interface{
	generator() string
}

type sample2 interface{
	function() string
}

func (c custom)generator()string{
	str := "hello world, "
	str += fmt.Sprintf("%v", c)
	return str
}

func function(i interface{}) {  
    fmt.Printf(fmt.Sprintf("%v\n",i))
}

func main() {
	var s sample
	fmt.Println(s)

	s = custom("go!")
	fmt.Println(s.generator())

	function("interface")
	
	t := struct {
        name string
    }{
        name: "name",
    }
	function(t)
}