package app

import (
	"context"
	"time"

	"tflgame"
	"tflgame/server/lib/cher"

	"github.com/rickb777/date/period"
)

func (a *App) GetGameState(ctx context.Context, req *tflgame.GetGameStateRequest) (*tflgame.GetGameStateResponse, error) {
	game, err := a.db.Q.GetGame(ctx, req.GameID)
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

	score, _, err := a.CalculateGameScore(game.DifficultyOptions, prompts)
	if err != nil {
		return nil, err
	}

	var inProgress = true
	var p period.Period

	if game.FinishedAt != nil {
		inProgress = false
		duration := game.FinishedAt.Sub(game.CreatedAt)
		p, _ = period.NewOf(duration)
	}

	bgCtx, can := context.WithTimeout(context.Background(), 1*time.Minute)
	go a.HandleEndgameEvents(bgCtx, can, req.GameID)

	return &tflgame.GetGameStateResponse{
		InProgress:        inProgress,
		Score:             score,
		GameTime:          p.String(),
		DifficultyOptions: game.DifficultyOptions,
		GameOptions:       game.GameOptions,
	}, nil
}
