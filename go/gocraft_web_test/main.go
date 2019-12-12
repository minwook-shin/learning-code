package main

import (
	"fmt"
	"github.com/gocraft/web"
	"net/http"
	"strings"
)

type context struct {
	Num int
}

func (c *context) testNum(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.Num = 10
	next(rw, req)
}

func (c *context) print(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.Num))
}

func main() {
	router := web.New(context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware((*context).testNum).
		Get("/", (*context).print)

	http.ListenAndServe("localhost:8080", router)
}
