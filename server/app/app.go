package app

import (
	"tflgame/server/db"
	"tflgame/server/lib/config"
	"tflgame/server/lib/tfl"

	"github.com/sirupsen/logrus"
)

// App is the main struct to attach business logic to
type App struct {
	db          *db.DB
	tfl         *tfl.Client
	SigningKeys *config.SigningKey

	Logger *logrus.Logger
}

func New(database *db.DB, sk *config.SigningKey, tfl *tfl.Client, logger *logrus.Logger) *App {
	return &App{
		db:          database,
		tfl:         tfl,
		SigningKeys: sk,
		Logger:      logger,
	}
}
