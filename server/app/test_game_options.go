package app

import (
	"context"

	"tflgame"
)

func (a *App) TestGameOptions(ctx context.Context, req *tflgame.GameOptions) (*tflgame.TestGameOptionsResponse, error) {
	promptOptions, err := a.db.Q.GetPossiblePrompts(ctx, req, 0)
	if err != nil {
		return nil, err
	}

	return &tflgame.TestGameOptionsResponse{
		PossiblePrompts: len(promptOptions),
	}, nil
}
