package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow)

	sec := timeNow.Unix()
	fmt.Println(sec)

    nano := timeNow.UnixNano()
	fmt.Println(nano)

    millis := nano / 1000000
    fmt.Println(millis)
	
    fmt.Println(time.Unix(sec, 0))
    fmt.Println(time.Unix(0, nano))
}
