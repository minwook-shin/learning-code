package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestFunc(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://example.com/test",
		httpmock.NewStringResponder(200, `[{"hello": "world"}]`))

	httpmock.RegisterResponder("GET", `=~^http://example.com/test/id/\d+\z`,
		httpmock.NewStringResponder(200, `{"id": 1, "hello": "world"}`))

	res, _ := http.Get("http://example.com/test")
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	res, _ = http.Get("http://example.com/test/id/1")
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	httpmock.GetTotalCallCount()

	countInfo := httpmock.GetCallCountInfo()
	match := countInfo["GET http://example.com/test"]
	fmt.Println(match)

	regexMatch := countInfo[`GET =~^http://example.com/test/id/\d+\z`]
	fmt.Println(regexMatch)
}
