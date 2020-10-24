package app

import (
	"context"

	"tflgame"
	"tflgame/server/db"
	"tflgame/server/lib/cher"
)

func (a *App) ReleaseHandle(ctx context.Context, req *tflgame.ReleaseHandleRequest) error {
	return a.db.DoTx(ctx, func(qw *db.QueryableWrapper) error {
		currentGame, err := qw.GetCurrentGame(ctx, req.UserID)
		if err != nil {
			v, ok := err.(cher.E)
			if !ok || v.Code != cher.NotFound {
				return err
			}
		}

		if currentGame != "" {
			return cher.New("game_in_progress", cher.M{
				"game_id": currentGame,
			})
		}

		return a.db.ReleaseHandle(ctx, req.UserID)
	})
}
