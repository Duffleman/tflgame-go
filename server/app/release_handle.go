package app

import (
	"context"

	"tflgame"
)

func (a *App) ReleaseHandle(ctx context.Context, req *tflgame.ReleaseHandleRequest) error {
	return a.db.ReleaseHandle(ctx, req.UserID)
}
