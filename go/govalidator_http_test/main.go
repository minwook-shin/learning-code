package main

import (
	"encoding/json"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"required_text": []string{"required", "between:4,8"},
		"email":         []string{"email"},
		"url":           []string{"url"},
		"phone":         []string{"digits:11"},
		"date":          []string{"date"},
		"boolean":       []string{"bool"},
	}

	messages := govalidator.MapData{
		"required_text": []string{"required:required TEXT"},
	}

	options := govalidator.Options{
		Request:         r,
		Rules:           rules,
		Messages:        messages,
		RequiredDefault: true,
	}

	validator := govalidator.New(options)
	values := validator.Validate()
	err := map[string]interface{}{"ERROR": values}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

// http://localhost:8000/?required_text=hello&url=www.google.com&email=test@example.com&phone=00000000000&date=2019-11-21&boolean=true
