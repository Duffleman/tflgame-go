package db

import (
	"context"
	"encoding/json"
	"time"

	"tflgame"

	"github.com/cuvva/ksuid-go"
)

func (qw *QueryableWrapper) GiveHint(ctx context.Context, userID, gameID string, payload *tflgame.GiveHintPayload) error {
	eventID := ksuid.Generate("event").String()
	now := time.Now().Format(time.RFC3339)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
			INSERT INTO events
			(id, type, user_id, game_id, payload, created_at)
			VALUES($1, $2, $3, $4, $5, $6)
		`,
		eventID, "give_hint", userID, gameID, payloadBytes, now,
	)
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
		UPDATE proj_prompts
		SET prompt = $2, hint_given_at = $3
		WHERE id = $1
	`, payload.PromptID, payload.NewPrompt, now)

	return err
}
