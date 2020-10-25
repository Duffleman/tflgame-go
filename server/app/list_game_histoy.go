package app

import (
	"context"

	"tflgame"
)

func (a *App) ListGameHistory(ctx context.Context, req *tflgame.ListGameHistoryRequest) ([]*tflgame.GameLog, error) {
	games, err := a.db.Q.ListAllGames(ctx, req.UserID, true)
	if err != nil {
		return nil, err
	}

	gameHistory := []*tflgame.GameLog{}

	for _, g := range games {
		gameHistory = append(gameHistory, &tflgame.GameLog{
			GameID: g.ID,
			Score:  g.Score,
		})
	}

	return gameHistory, nil
}
