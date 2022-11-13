package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	e := echo.New()
	fileLogger := &lumberjack.Logger{
		Filename:   "/var/log/efk/request.log",
		MaxSize:    20, // megabytes
		MaxBackups: 3,
		MaxAge:     1, //days
		Compress:   true, // disabled by default
	}

	logger := zerolog.New(fileLogger)
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
