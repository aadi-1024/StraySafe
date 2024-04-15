package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

var app = &App{}

func main() {

	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	app.Db = db
	app.JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	e := echo.New()
	SetupRoutes(e)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
