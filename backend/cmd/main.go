package main

import (
	"log"
	"net/http"

	"github.com/aadi-1024/straysafe/internals/db"
	"github.com/aadi-1024/straysafe/internals/mailer"
)

var App *Config

func main() {
	var err error
	App = &Config{}

	App.Mail = mailer.NewMailer("username", "password", "localhost")
	go App.Mail.StartService()
	App.Db, err = db.NewDb("mongodb://localhost:27017")

	if err != nil {
		log.Println(err)
	}

	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: NewRouter(),
	}

	log.Println("starting server on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
