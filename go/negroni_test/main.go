package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello, world!")
	})

	negroniObj := negroni.Classic()

	negroniObj.UseHandler(serveMux)

	http.ListenAndServe(":8080", negroniObj)
}
