package main

import (
	"fmt"
)

func main() {
	type example1 struct{
		name string
		age int
	}
	e1 := example1{
		name : "name",
		age : 21,
	}
	fmt.Println(e1)

	type example2 struct{
		firstName, lastName string
		age int
	}
	e2 := example2{
		firstName : "first",
		lastName : "last",
		age: 21,
	}
	fmt.Println(e2)

	e3 := struct{
		name string
		age int
		}{
		name :"name",
		age : 21,
	}
	fmt.Println(e3)

	
	e4 := struct{
		name string
		age int
		}{
		name :"name",
		age : 21,
	}
	fmt.Println(e4.name,e4.age)

	type example5 struct {  
		firstName, lastName string
		age        int
	}
	e5 := &example5{"first","last",21}
	fmt.Println(e5,*e5,(*e5).age,e5.age)
	
	type example6 struct {  
		string
		int
	}
	e6 := example6{"first",21}
	fmt.Println(e6)

	type nested struct{
		hello string
	}
	type structs struct{
		world nested
	}
	var e7 structs
	e7.world = nested{
		hello : "hello, wolrd!",
	}
	fmt.Println(e7)

	type equals struct{
		text string
	}
	e8 := equals{
		text : "original",
	}
	e8Copy := equals{
		text : "copy",
	}
	fmt.Println(e8 == e8Copy)
}