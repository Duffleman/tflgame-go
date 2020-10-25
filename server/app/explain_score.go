package app

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
)

func (a *App) ExplainScore(ctx context.Context, req *tflgame.ExplainScoreRequest) (*tflgame.Calculations, error) {
	if req.GameID == nil {
		_, calc, err := a.CalculateUserScore(ctx, nil, req.UserID)
		if err != nil {
			return nil, err
		}

		return calc, nil
	}

	game, err := a.db.Q.GetGame(ctx, *req.GameID)
	if err != nil {
		return nil, err
	}

	if game.UserID != req.UserID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	prompts, err := a.db.Q.ListPrompts(ctx, *req.GameID)
	if err != nil {
		return nil, err
	}

	_, calcs, err := a.CalculateGameScore(game.DifficultyOptions, prompts)
	if err != nil {
		return nil, err
	}

	return calcs, nil
}
