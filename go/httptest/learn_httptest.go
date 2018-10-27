package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	request := httptest.NewRequest("GET","http://example.com", nil)
	responseRecorder := httptest.NewRecorder()
	handler(responseRecorder, request)

	response := responseRecorder.Result()
	bodyByte, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(http.StatusText(response.StatusCode))
	fmt.Println(string(bodyByte))
}