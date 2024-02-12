package main

import (
	"github.com/aadi-1024/straysafe/internals/db"
	"github.com/aadi-1024/straysafe/internals/mailer"
)

type Config struct {
	Mail *mailer.Mailer
	Db *db.Database
}