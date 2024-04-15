package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/handlers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
	e.POST("/incident", handlers.IncidentPostHandler(app.Db), JwtMiddleware) //new incident
	e.POST("/resolveIncident", handlers.MarkResolvedHandler(app.Db), func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("typ", "ngo")
			c.Set("id", 1)
			return next(c)
		}
	})
}
