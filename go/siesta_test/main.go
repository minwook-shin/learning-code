package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VividCortex/siesta"
)

func main() {
	service := siesta.NewService("/")

	get := "GET"
	service.Route(get, "/", "'Hello, world!'",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, world!")
		})

	timeHandler := func(c siesta.Context, w http.ResponseWriter, r *http.Request) {
		timeDelta := time.Now().Sub(c.Get("time").(time.Time))
		fmt.Fprintf(w, "%v", timeDelta)
	}

	contextHandler := siesta.Compose(func(c siesta.Context, w http.ResponseWriter, r *http.Request) {
		c.Set("time", time.Now())
	}, timeHandler)

	service.Route("GET", "/time", "usage", contextHandler)

	err := http.ListenAndServe(":8080", service)
	if err != nil {
		log.Fatal(err)
	}
}
