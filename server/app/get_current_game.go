package app

import (
	"context"

	"tflgame"
)

func (a *App) GetCurrentGame(ctx context.Context, req *tflgame.GetCurrentGameRequest) (*tflgame.GetCurrentGameResponse, error) {
	gameID, err := a.db.Q.GetCurrentGame(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	prompt, err := a.db.Q.GetNextPrompt(ctx, gameID)
	if err != nil {
		return nil, err
	}

	var next *tflgame.NextPrompt

	if prompt != nil {
		next = &tflgame.NextPrompt{
			PromptID: prompt.ID,
			Prompt:   prompt.Prompt,
			Length:   prompt.Length,
		}
	}

	return &tflgame.GetCurrentGameResponse{
		GameID: gameID,
		Next:   next,
	}, nil
}
