package main

import "fmt"

func insertNum() func(n int) int {
    total := 0
    return func(n int) int {
        total = total + n
        return total
    }
}

func main() {

    adder := insertNum()

    fmt.Println(adder(4))
    fmt.Println(adder(400))
}