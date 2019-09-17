package main

import (
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func handlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	svg := svg.New(w)
	width := 1000
	height := 1000
	svg.Start(width, height)
	svg.Circle(width/2, height/2, 200, "style")
	svg.Text(width/2, height/2, "Hello, World!", "text-anchor:middle;font-size:50px;fill:white")
	svg.End()
}

func main() {
	http.Handle("/test", http.HandlerFunc(handlerFunc))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
