package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"hello, world!\"}"))
	})

	handler := cors.Default().Handler(serveMux)

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080", "http://example:8080"},
		AllowedMethods:   []string{http.MethodGet},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           0,
		Debug:            true,
	})

	handler = corHandler.Handler(handler)
	http.ListenAndServe(":8080", handler)
}
