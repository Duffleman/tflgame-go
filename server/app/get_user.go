package app

import (
	"context"

	"tflgame"
)

func (a *App) GetUser(ctx context.Context, userID string) (*tflgame.User, error) {
	return a.db.Q.GetUserByID(ctx, userID)
}
