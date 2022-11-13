package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	e := echo.New()
	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().Str("URI", v.URI).Int("status", v.Status).Time("time", v.StartTime.Local()).Msg("Request")

			return nil
		},
	}))

	e.GET("/", handler)
	logger.Fatal().Err(e.Start(":8080")).Msg("Server started")
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
