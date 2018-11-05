package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow)

	fmt.Println(timeNow.Location())
	fmt.Println(timeNow.Weekday())

	timeCustom := time.Date(2017, 12, 22, 12, 00, 00, 000000000, time.FixedZone("Asia/Seoul", 9*60*60))
	fmt.Println(timeCustom)

	fmt.Println(timeCustom.Location())
	fmt.Println(timeCustom.Weekday())

	fmt.Println(timeCustom.Equal(timeNow))

	diff := timeNow.Sub(timeCustom)
	fmt.Println(diff)

	fmt.Println(timeCustom.Add(diff))
	fmt.Println(timeCustom.Add(diff).Equal(timeNow))
}
