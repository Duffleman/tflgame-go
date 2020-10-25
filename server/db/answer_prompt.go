package db

import (
	"context"
	"encoding/json"
	"time"

	"tflgame"

	"github.com/cuvva/ksuid-go"
)

// AnswerPrompt answers a given prompt. You **must** run this within a transaction.
func (qw *QueryableWrapper) AnswerPrompt(ctx context.Context, prompt *tflgame.Prompt) error {
	eventID := ksuid.Generate("event").String()

	payloadBytes, err := json.Marshal(tflgame.AnswerPromptPayload{
		PromptID:    prompt.ID,
		AnswerGiven: *prompt.AnswerGiven,
		Correct:     prompt.Correct,
	})
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
			INSERT INTO events
			(id, type, user_id, game_id, payload, created_at)
			VALUES($1, $2, $3, $4, $5, $6)
		`,
		eventID, "answer_prompt", prompt.UserID, prompt.GameID, payloadBytes, prompt.AnsweredAt.Format(time.RFC3339),
	)
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
		UPDATE proj_prompts
		SET correct = $2, answer_given = $3, answered_at = $4
		WHERE id = $1
	`, prompt.ID, prompt.Correct, prompt.AnswerGiven, prompt.AnsweredAt.Format(time.RFC3339))

	return err
}
