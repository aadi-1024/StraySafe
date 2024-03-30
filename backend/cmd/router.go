package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})

}
