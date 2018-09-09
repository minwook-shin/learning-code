package main

import (
	"fmt"
)

func helloworld(){}

func printValue(){
	fmt.Println(3.14)
}

func sum(a,b int){
	fmt.Println(a + b)
}

func returnValue()float64{
	return 3.14
}

func returnSum(a,b int)int{
	return a + b
}

func returnTwoValue(age,count int)(int,int){
	value1 := age
	value2 := count + 1 
	return value1, value2
}

func named(a,b int)(r1 , r2 int){
	r1 = a + 1
	r2 = b + 1
	return
}

func main() {
	helloworld()

	printValue()
	
	fmt.Println(returnValue())
	
	sum(1,2)
	
	fmt.Println(returnSum(1,2))
	
	fmt.Println(returnTwoValue(20,0))
	
	fmt.Println(named(1,2))
	
	var tmp, _ = returnTwoValue(1,2)
	fmt.Println(tmp)
}