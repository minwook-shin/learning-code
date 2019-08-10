package main

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("setting.ini")

	if err != nil {
		panic(err)
	}
	Bool, err := cfg.Section("types").Key("Bool").Bool()
	String := cfg.Section("types").Key("String").String()
	Int, err := cfg.Section("types").Key("Int").Int()
	Uint, err := cfg.Section("types").Key("Uint").Uint()
	TimeFormat, err := cfg.Section("types").Key("TimeFormat").TimeFormat(time.RFC3339)
	Time, err := cfg.Section("types").Key("Time").Time()

	MustString := cfg.Section("Must").Key("MustString").MustString("default")
	MustBool := cfg.Section("Must").Key("MustBool").MustBool(true)
	MustInt := cfg.Section("Must").Key("MustInt").MustInt(10)
	MustUint := cfg.Section("Must").Key("MustUint").MustUint(10)
	MustTimeFormat := cfg.Section("Must").Key("MustTimeFormat").MustTimeFormat(time.RFC3339, time.Now())
	MustTime := cfg.Section("Must").Key("MustTime").MustTime(time.Now())

	In := cfg.Section("InTypes").Key("In").In("default", []string{"one", "two", "three"})
	InInt := cfg.Section("InTypes").Key("InInt").InInt(0, []int{10, 20, 30})

	RangeInt := cfg.Section("Range").Key("RangeInt").RangeInt(1, 0, 20)

	Strings := cfg.Section("delimiter").Key("Strings").Strings(",")
	Ints := cfg.Section("delimiter").Key("Ints").Ints(",")
	Uints := cfg.Section("delimiter").Key("Uints").Uints(",")
	Times := cfg.Section("delimiter").Key("Times").Times(",")

	ValidInts := cfg.Section("Valid").Key("ValidInts").ValidInts(",")
	ValidTimes := cfg.Section("Valid").Key("ValidTimes").ValidTimes(",")

	StrictInts, err := cfg.Section("Strict").Key("INTS").StrictInts(",")

	fmt.Print(Bool, String, Int, Uint, TimeFormat, Time,
		MustString, MustBool, MustInt, MustUint, MustTimeFormat, MustTime,
		In, InInt, RangeInt, Strings, Ints, Uints, Times,
		ValidInts, ValidTimes, StrictInts)

	cfg.Section("").Key("Write_Value").SetValue("SetValue")
	cfg.SaveTo("setiing.ini.local")
}
