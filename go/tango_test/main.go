package main

import (
	"github.com/lunny/tango"
)

type jsonAction struct {
	tango.JSON
}

func (jsonAction) Get() interface{} {
	return map[string]string{
		"message": "Hello world!",
	}
}

func main() {
	tango := tango.Classic()

	tango.Get("/", new(jsonAction))

	tango.Server.Addr = "0.0.0.0:8000"
	tango.Server.Handler = tango

	err := tango.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
