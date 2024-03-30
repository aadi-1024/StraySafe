package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/postgres"
)

// App carries the app wide config and units
type App struct {
	Db *postgres.Database
}
