package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

var app = &App{}

func main() {

	db, err := postgres.NewDb(os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
	app.Db = db

	e := echo.New()
	SetupRoutes(e)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
