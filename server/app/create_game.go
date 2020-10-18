package app

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
)

func (a *App) CreateGame(ctx context.Context, req *tflgame.CreateGameRequest) (*tflgame.CreateGameResponse, error) {
	promptsStr, err := a.db.Q.GetPossiblePrompts(ctx, &req.GameOptions)
	if err != nil {
		return nil, err
	}

	if len(promptsStr) < req.DifficultyOptions.Rounds {
		return nil, cher.New("round_mismatch", cher.M{
			"possible_prompts": len(promptsStr),
			"requested_rounds": req.DifficultyOptions.Rounds,
		})
	}

	prompts := a.GeneratePrompts(promptsStr, req.DifficultyOptions)

	gameID, err := a.db.CreateGame(ctx, req, prompts)
	if err != nil {
		return nil, err
	}

	return &tflgame.CreateGameResponse{
		ID:       gameID,
		PromptID: prompts[0].ID,
		Prompt:   prompts[0].Prompt,
		Length:   prompts[0].Length,
	}, nil
}
