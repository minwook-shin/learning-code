package main

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func test() {
	testVal := 1
	spew.Dump(testVal)

	type testStruct struct {
		ID   int64  `json:"project_id"`
		Name string `json:"name"`
	}

	t := testStruct{ID: 1, Name: "TEST"}
	spew.Dump(t)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hello World!")
	fmt.Println(spew.Sdump(w))
}

func main() {
	test()
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
}
