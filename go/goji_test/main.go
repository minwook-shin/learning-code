package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func handler(w http.ResponseWriter, r *http.Request) {
	parameter := pat.Param(r, "q")
	fmt.Fprintf(w, parameter)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/search/:q"), handler)

	http.ListenAndServe("localhost:8080", mux)
}
