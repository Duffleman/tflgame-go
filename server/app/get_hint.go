package app

import (
	"context"
	"strings"

	"tflgame"
	"tflgame/server/db"
	"tflgame/server/lib/cher"
)

var hintRemove = []string{"A", "E", "I", "O", "U"}

func (a *App) GetHint(ctx context.Context, req *tflgame.GetHintRequest) (*tflgame.GetHintResponse, error) {
	var lines []string
	var prompt *tflgame.Prompt
	var newPrompt string
	var err error

	err = a.db.DoTx(ctx, func(qw *db.QueryableWrapper) error {
		prompt, err = qw.GetPrompt(ctx, req.PromptID)
		if err != nil {
			return err
		}

		if prompt.AnsweredAt != nil {
			return cher.New("already_answered", nil)
		}

		lines, err = a.db.Q.GetAllLines(ctx, prompt.Answer)
		if err != nil {
			return err
		}

		newPrompt = prompt.Answer

		for _, v := range alwaysRemove {
			newPrompt = strings.Replace(newPrompt, v, "", -1)
		}

		if prompt.HintGivenAt == nil {
			err := qw.GiveHint(ctx, prompt.UserID, prompt.GameID, &tflgame.GiveHintPayload{
				PromptID:  prompt.ID,
				Lines:     lines,
				NewPrompt: newPrompt,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &tflgame.GetHintResponse{
		Prompt: newPrompt,
		Lines:  lines,
	}, nil
}
