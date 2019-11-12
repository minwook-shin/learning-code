package main

import (
	"fmt"
	"time"

	"github.com/beefsack/go-rate"
)

var rateLimiter = rate.New(3, time.Second)

func main() {

	now := time.Now()

	for i := 1; i <= 5; i++ {
		rateLimiter.Wait()
		fmt.Println(i, time.Now().Sub(now))
	}

	rateLimiter1 := rate.New(1, time.Second)
	rateLimiter2 := rate.New(2, time.Second*3)

	for i := 1; i <= 5; i++ {
		rateLimiter1.Wait()
		rateLimiter2.Wait()
		fmt.Println(i, time.Now().Sub(now))
	}

	for i := 1; i <= 5; i++ {
		if ok, remaining := rateLimiter.Try(); ok {
			fmt.Println(fmt.Sprintln(i))
		} else {
			fmt.Println(remaining)
		}
	}

	time.Sleep(time.Second / 2)
	if ok, remaining := rateLimiter.Try(); ok {
		fmt.Println("Okay1")
	} else {
		fmt.Println(remaining)
	}
}
