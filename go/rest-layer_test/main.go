package main

import (
	"log"
	"net/http"

	"github.com/rs/rest-layer/resource"
	"github.com/rs/rest-layer/resource/testing/mem"
	"github.com/rs/rest-layer/rest"
	"github.com/rs/rest-layer/schema"
)

var (
	test = schema.Schema{
		Description: `test object`,
		Fields: schema.Fields{
			"id": {
				Required:   true,
				ReadOnly:   true,
				Filterable: true,
				Sortable:   true,
				OnInit:     schema.NewID,
				Validator: &schema.String{
					Regexp: "^[0-9a-z]$",
				},
			},
			"created": {
				Required:   true,
				ReadOnly:   true,
				Filterable: true,
				Sortable:   true,
				OnInit:     schema.Now,
				Validator:  &schema.Time{},
			},
		},
	}
)

func main() {
	index := resource.NewIndex()

	index.Bind("tests", test, mem.NewHandler(), resource.Conf{AllowedModes: resource.ReadWrite})

	api, _ := rest.NewHandler(index)

	http.Handle("/api/", http.StripPrefix("/api/", api))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
