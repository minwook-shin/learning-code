package main

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

func handler(ctx *macaron.Context) string {
	return "hello, world!"
}

func main() {
	classicMacaron := macaron.Classic()

	classicMacaron.Get("/", func() string {
		return "root"
	})

	classicMacaron.Get("/handler", handler)

	// classicMacaron.Run()

	err := http.ListenAndServe("0.0.0.0:4000", classicMacaron)
	if err != nil {
		panic(err)
	}
}
