package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := fasthttp.CompressHandler(requestHandler)
	if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
		panic(err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "%s", &ctx.Request)
	fmt.Fprintf(ctx, "%s", &ctx.Response)

	ctx.SetContentType("text/plain; charset=utf8")

	ctx.Response.Header.Set("X-Custom-Header", "Custom-header-value")

	var cookie fasthttp.Cookie
	cookie.SetKey("cookie-name")
	cookie.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&cookie)
}
