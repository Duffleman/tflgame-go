package app

import (
	"tflgame/server/db"
	"tflgame/server/lib/config"
)

// App is the main struct to attach business logic to
type App struct {
	db          *db.DB
	SigningKeys *config.SigningKey
}

func New(database *db.DB, sk *config.SigningKey) *App {
	return &App{
		db:          database,
		SigningKeys: sk,
	}
}
