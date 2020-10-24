package db

import (
	"context"
	"encoding/json"
	"time"

	"tflgame"

	"github.com/cuvva/ksuid-go"
)

// RecalculateGameScore recalculates the score for a game. You **must** run this within a transaction.
func (qw *QueryableWrapper) RecalculateGameScore(ctx context.Context, userID, gameID string, score int) error {
	eventID := ksuid.Generate("event").String()
	now := time.Now()

	payloadBytes, err := json.Marshal(&tflgame.RecalculateGameScorePayload{
		Score: score,
	})
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
			INSERT INTO events
			(id, type, user_id, game_id, payload, created_at)
			VALUES($1, $2, $3, $4, $5, $6)
		`,
		eventID, "recalculate_game_score", userID, gameID, payloadBytes, now.Format(time.RFC3339),
	)
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
		UPDATE proj_games
		SET score = $2
		WHERE id = $1
	`, gameID, score)

	return err
}
