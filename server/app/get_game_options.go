package app

import (
	"context"

	"tflgame"
)

func (a *App) GetGameOptions(ctx context.Context) (*tflgame.GetGameOptionsResponse, error) {
	return a.db.Q.GetGameOptions(ctx)
}
