package app

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
)

func (a *App) ExplainScore(ctx context.Context, req *tflgame.ExplainScoreRequest) (*tflgame.Calculations, error) {
	game, err := a.db.Q.GetGameOpts(ctx, req.GameID)
	if err != nil {
		return nil, err
	}

	if game.UserID != req.UserID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	prompts, err := a.db.Q.ListPrompts(ctx, req.GameID)
	if err != nil {
		return nil, err
	}

	_, calcs, err := a.CalculateScore(game.DifficultyOptions, prompts)
	if err != nil {
		return nil, err
	}

	return calcs, nil
}
