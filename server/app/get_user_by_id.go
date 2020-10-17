package app

import (
	"context"

	"tflgame"
)

func (a *App) GetUserByID(ctx context.Context, userID string) (*tflgame.User, error) {
	return a.db.Q.GetUserByID(ctx, userID)
}
