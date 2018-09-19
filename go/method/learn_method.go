package main

import (
	"fmt"
)

type sample struct{
	str string
	i int
}

type mathPI struct{
	pi float32
}

func (e sample)method()  {
	fmt.Println(e.str,e.i)
}

func method2(e sample)  {
	fmt.Println(e.str,e.i)
}

func (m mathPI)method3() float32{
	return m.pi
}


type ptr struct{
	i int
}
func (p *ptr)method4(new int){
	p.i = new
}

func main() {
	example := sample{
		"hello, world!",
		10,
	}
	example.method()

	example2 := sample{
		"hello, world!",
		10,
	}
	method2(example2)

	m := mathPI{
		3.14,
	}
	fmt.Println(m.method3())

	p := ptr{
		i : 10,
	}
	fmt.Println(p.i)
	(&p).method4(20)
	fmt.Println(p.i)
	fmt.Println(p.i)
	p.method4(30)
	fmt.Println(p.i)
}