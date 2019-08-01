package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	str  = kingpin.Flag("str", "string description").String()
	bool = kingpin.Flag("bool", "bool description").Bool()
	// duration = kingpin.Flag("duration", "duration description").Required().Short('t').Duration()
	ip     = kingpin.Arg("ip", "IP description").Required().IP()
	url    = kingpin.Arg("url", "url description").Required().URL()
	number = kingpin.Arg("int", "int description").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
}
