package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

func main() {
	renderObject := render.New()
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello, world!"))
	})

	serveMux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		renderObject.Data(w, http.StatusOK, []byte("hello, world!"))
	})

	serveMux.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
		renderObject.Text(w, http.StatusOK, "hello, text!")
	})

	serveMux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		renderObject.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
	})

	serveMux.HandleFunc("/jsonp", func(w http.ResponseWriter, req *http.Request) {
		renderObject.JSONP(w, http.StatusOK, "callback", map[string]string{"hello": "jsonp"})
	})

	serveMux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		renderObject.HTML(w, http.StatusOK, "test", "text")
	})

	http.ListenAndServe("127.0.0.1:8080", serveMux)
}
