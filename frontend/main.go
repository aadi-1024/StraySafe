package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()
	e.Static("/", ".")
	if err := e.Start("0.0.0.0:80"); err != nil {
		log.Fatalln(err)
	}
}