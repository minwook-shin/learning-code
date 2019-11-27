package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler)
	err := e.Start(":8000")
	if err != nil {
		e.Logger.Fatal(err)
	}
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
