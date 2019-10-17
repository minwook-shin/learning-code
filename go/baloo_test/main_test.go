package test

import (
	"errors"
	"net/http"
	"testing"

	"gopkg.in/h2non/baloo.v3"
)

var balooClient = baloo.New("http://httpbin.org")

const schema = `{
	"title": "Example Schema",
	"type": "object",
	"properties": {
	  "origin": {
		"type": "string"
	  }
	},
	"required": ["origin"]
  }`

func assert(res *http.Response, req *http.Request) error {
	if res.StatusCode >= 400 {
		return errors.New("Invalid response")
	}
	return nil
}

func TestBFunc(t *testing.T) {
	balooClient.Get("/get").
		Expect(t).
		Status(200).
		Header("Server", "nginx").
		Type("json").
		JSONSchema(schema).
		AssertFunc(assert).
		Done()
}
