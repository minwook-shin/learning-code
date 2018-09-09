package calc

import "fmt"

func init(){
	fmt.Println("lib init")
}

func Sum(a,b float64) float64 {
	return a + b
}

func Sub(a,b float64) float64 {
	return a - b
}

func Div(a,b float64) float64 {
	return a / b
}

func Mul(a,b float64) float64 {
	return a * b
}