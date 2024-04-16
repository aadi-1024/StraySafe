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
	e.GET("/resolveIncident/:id", handlers.MarkResolvedHandler(app.Db), JwtMiddleware)
	e.GET("/ngo/about/:id", handlers.GetNgoByIdHandler(app.Db))
	e.POST("/ngo/nearest", handlers.NearestNgos(app.Db))
}
