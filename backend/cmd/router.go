package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
	e.POST("/user/incident", handlers.IncidentPostHandler(app.Db), JwtMiddleware) //new incident
	e.GET("/ngo/resolveIncident/:id", handlers.MarkResolvedHandler(app.Db), JwtMiddleware)
	e.GET("/ngo/about/:id", handlers.GetNgoByIdHandler(app.Db))
	e.POST("/ngo/nearest", handlers.NearestNgos(app.Db))
	e.GET("/ngo/dashboard", handlers.GetDashboardHandler(app.Db), JwtMiddleware)
	e.GET("/ngo/incident/:id", handlers.GetIncidentHandler(app.Db), JwtMiddleware)
	//e.GET("/ngo/incident/img/:filename", handlers.GetImage(), JwtMiddleware)
	e.Static("static/", "img/")
}
