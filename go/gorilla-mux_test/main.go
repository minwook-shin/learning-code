package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world!\n"))
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, variables["id"])
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler)
	router.HandleFunc("/user/{id}", paramHandler)
	router.HandleFunc("/regex/{id:[0-9]+}", paramHandler)

	router.Use(logging)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
