package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handler)
	e.Logger.Fatal(e.Start(":1323"))
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
