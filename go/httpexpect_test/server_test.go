package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

type userMap map[string]interface{}

func handlerFunc() http.Handler {
	mux := http.NewServeMux()

	user := userMap{}
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			v := []string{}
			for u := range user {
				v = append(v, u)
			}
			byteData, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(byteData)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	return mux
}

func TestFunc(t *testing.T) {
	handler := handlerFunc()

	server := httptest.NewServer(handler)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	expect.GET("/user").
		Expect().
		Status(http.StatusOK).JSON().Array().Empty()

	expect.GET("/user/test_name").
		Expect().
		Status(http.StatusNotFound)
}
