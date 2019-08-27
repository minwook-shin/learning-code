package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow)

	timeNow = now.New(time.Date(2000, 1, 1, 00, 00, 00, 0, time.Now().Location())).Local()
	fmt.Println(timeNow)

	timeNow = now.BeginningOfMinute()
	fmt.Println(timeNow)
	timeNow = now.BeginningOfHour()
	fmt.Println(timeNow)
	timeNow = now.BeginningOfDay()
	fmt.Println(timeNow)
	timeNow = now.BeginningOfWeek()
	fmt.Println(timeNow)
	timeNow = now.BeginningOfMonth()
	fmt.Println(timeNow)
	timeNow = now.BeginningOfQuarter()
	fmt.Println(timeNow)
	timeNow = now.BeginningOfYear()
	fmt.Println(timeNow)

	timeNow = now.EndOfMinute()
	fmt.Println(timeNow)
	timeNow = now.EndOfHour()
	fmt.Println(timeNow)
	timeNow = now.EndOfDay()
	fmt.Println(timeNow)
	timeNow = now.EndOfWeek()
	fmt.Println(timeNow)
	timeNow = now.EndOfMonth()
	fmt.Println(timeNow)
	timeNow = now.EndOfQuarter()
	fmt.Println(timeNow)
	timeNow = now.EndOfYear()
	fmt.Println(timeNow)

	timeParser, err := now.Parse("2019")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeParser)
	timeParser, err = now.Parse("2019-08")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeParser)
	timeParser, err = now.Parse("2019-08-27")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeParser)
	timeParser, err = now.Parse("2019-08-27 12:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeParser)
	timeParser, err = now.Parse("2019-08-27 12:00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeParser)
	timeParser, err = now.Parse("12:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeParser)

	timeParser = now.MustParse("08-27")
	fmt.Println(timeParser)
	timeParser = now.MustParse("2019-08-27")
	fmt.Println(timeParser)
}
