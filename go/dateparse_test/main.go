package main

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func main() {
	timeVar, err := dateparse.ParseAny("2019/08/26")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)

	// timeVar = dateparse.MustParse("hello")

	timeVar, err = dateparse.ParseStrict("2019/08/26")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)

	layout, err := dateparse.ParseFormat("Aug 8, 2019 00:00:00 PM")
	if err != nil {
		panic(err)
	}
	fmt.Println(layout)

	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}

	timeVar, err = dateparse.ParseIn("2019/08/26", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)

	time.Local = loc
	timeVar, err = dateparse.ParseLocal("Aug 26, 2019 00:00:00 PM")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("Aug 26, 2019")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("Mon Aug 26 00:00:00 2019")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("August 26, 2019 00:00am")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("August 26, 2019, 00:00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("August 26, 2019")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("August 26th, 2019")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("26 Aug 2019, 00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("26 Aug 2019 00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("26 Aug 2019")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("26 August 2019")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("2019-Aug-26")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("2019/08/26")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("2019/08/26 00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("2019/08/26 00:00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("2019-08-26")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("20190626")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
	timeVar, err = dateparse.ParseLocal("20190826000000")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeVar)
}
