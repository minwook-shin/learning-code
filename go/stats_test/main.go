package main

import (
	"encoding/json"
	"github.com/thoas/stats"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware := stats.New()
		stats := middleware.Data()

		byteStats, err := json.Marshal(stats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(byteStats)
	})

	handler := stats.New().Handler(h)

	http.ListenAndServe(":8080", handler)
}
