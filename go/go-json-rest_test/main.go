package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	restAPI := rest.NewApi()

	restAPI.Use(rest.DefaultDevStack...)

	restAPI.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))

	http.Handle("/", http.StripPrefix("/", restAPI.MakeHandler()))

	restAPI2 := rest.NewApi()

	restAPI2.Use(rest.DefaultDevStack...)

	restApp, err := rest.MakeRouter(
		rest.Get("/param/#name", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": req.PathParam("name")})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	restAPI2.SetApp(restApp)
	http.Handle("/api/", http.StripPrefix("/api", restAPI2.MakeHandler()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
