package db

import (
	"context"
	"encoding/json"
	"time"

	"tflgame"

	"github.com/cuvva/ksuid-go"
)

// RecalculateUserScore recalculates the score for a user. You **must** run this within a transaction.
func (qw *QueryableWrapper) RecalculateUserScore(ctx context.Context, userID string, score int) error {
	eventID := ksuid.Generate("event").String()
	now := time.Now()

	payloadBytes, err := json.Marshal(&tflgame.RecalculateUserScorePayload{
		Score: score,
	})
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
			INSERT INTO events
			(id, type, user_id, payload, created_at)
			VALUES($1, $2, $3, $4, $5)
		`,
		eventID, "recalculate_user_score", userID, payloadBytes, now.Format(time.RFC3339),
	)
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
		UPDATE proj_users
		SET score = $2
		WHERE id = $1
	`, userID, score)

	return err
}
