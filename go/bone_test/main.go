package main

import (
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	num := bone.GetValue(req, "num")
	text := bone.GetValue(req, "text")

	rw.Write([]byte("hello, world! " + num + text))
}

func main() {
	mux := bone.New()

	mux.RegisterValidatorFunc("isNum", func(s string) bool {
		if _, err := strconv.Atoi(s); err == nil {
			return true
		}
		return false
	})

	mux.Get("/:num|isNum", http.HandlerFunc(handler))

	mux.Get("/:num/:text", http.HandlerFunc(handler))

	mux.Get("/#num^[0-9]$", http.HandlerFunc(handler))

	mux.Post("/post", http.HandlerFunc(handler))

	mux.Handle("/", http.HandlerFunc(handler))

	http.ListenAndServe(":8080", mux)
}
