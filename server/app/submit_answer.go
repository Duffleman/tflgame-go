package app

import (
	"context"
	"time"

	"tflgame"
	"tflgame/server/db"
	"tflgame/server/lib/cher"
)

func (a *App) SubmitAnswer(ctx context.Context, req *tflgame.SubmitAnswerRequest) (*tflgame.SubmitAnswerResponse, error) {
	var prompt, nextPrompt *tflgame.Prompt

	err := a.db.DoTx(ctx, func(qw *db.QueryableWrapper) error {
		var err error

		prompt, err = qw.GetPrompt(ctx, req.PromptID)
		if err != nil {
			return err
		}

		if prompt.AnsweredAt != nil {
			return cher.New("already_answered", cher.M{
				"prompt_id": prompt.ID,
				"game_id":   prompt.GameID,
			})
		}

		prompt.Correct = a.CompareAnswers(req.Answer, prompt.Answer)

		now := time.Now()
		prompt.AnsweredAt = &now

		prompt.AnswerGiven = &req.Answer

		err = qw.AnswerPrompt(ctx, prompt)
		if err != nil {
			return err
		}

		nextPrompt, err = qw.GetNextPrompt(ctx, prompt.GameID)
		if err != nil {
			v, ok := err.(cher.E)
			if !ok || v.Code != cher.NotFound {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := &tflgame.SubmitAnswerResponse{
		UserID:  prompt.UserID,
		Correct: prompt.Correct,
		Answer:  prompt.Answer,
	}

	if nextPrompt != nil {
		res.Next = &tflgame.NextPrompt{
			PromptID: nextPrompt.ID,
			Prompt:   nextPrompt.Prompt,
		}
	} else {
		bgCtx, can := context.WithTimeout(context.Background(), 1*time.Minute)
		go a.HandleEndgameEvents(bgCtx, can, prompt.GameID)
	}

	return res, nil
}
