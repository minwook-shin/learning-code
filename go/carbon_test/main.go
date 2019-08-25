package main

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	fmt.Println(carbon.Now().DateTimeString())

	today, _ := carbon.NowInLocation("Asia/Seoul")
	fmt.Println(today)

	carbonTime := today.AddYear()
	fmt.Println(carbonTime)
	carbonTime = today.AddYears(1)
	fmt.Println(carbonTime)

	carbonTime = today.AddMonth()
	fmt.Println(carbonTime)
	carbonTime = today.AddMonths(1)
	fmt.Println(carbonTime)

	carbonTime = today.AddWeekday()
	fmt.Println(carbonTime)
	carbonTime = today.AddWeekdays(1)
	fmt.Println(carbonTime)

	carbonTime = today.AddHour()
	fmt.Println(carbonTime)
	carbonTime = today.AddMinute()
	fmt.Println(carbonTime)
	carbonTime = today.AddSecond()
	fmt.Println(carbonTime)

	carbonTime = today.SubYear()
	fmt.Println(carbonTime)
	carbonTime = today.SubYears(1)
	fmt.Println(carbonTime)

	carbonTime = today.SubMonth()
	fmt.Println(carbonTime)
	carbonTime = today.SubMonths(1)
	fmt.Println(carbonTime)

	carbonTime = today.SubWeekday()
	fmt.Println(carbonTime)
	carbonTime = today.SubWeekdays(1)
	fmt.Println(carbonTime)

	carbonTime = today.SubHour()
	fmt.Println(carbonTime)
	carbonTime = today.SubMinute()
	fmt.Println(carbonTime)
	carbonTime = today.SubSecond()
	fmt.Println(carbonTime)

	testDate, _ := carbon.CreateFromDate(2019, 8, 25, "Asia/Seoul")

	fmt.Println(testDate.IsWeekday())
	fmt.Println(testDate.IsWeekend())

	fmt.Println(testDate.IsPast())
	fmt.Println(testDate.IsFuture())

	seoul, _ := carbon.CreateFromDate(2019, 8, 25, "Asia/Seoul")
	london, _ := carbon.CreateFromDate(2019, 8, 25, "Europe/London")
	fmt.Println(london.DiffInHours(seoul, true))

	parsed, _ := carbon.Parse(carbon.DateFormat, "2019-08-25", "Asia/Seoul")
	fmt.Println(parsed)

	EndOfWeekTime, _ := carbon.Create(2019, 8, 25, 12, 0, 0, 0, "Asia/Seoul")
	fmt.Println(EndOfWeekTime.EndOfWeek())

	nextWednesday, _ := carbon.Create(2019, 8, 25, 12, 0, 0, 0, "Asia/Seoul")
	fmt.Println(nextWednesday.Next(time.Wednesday))
}
