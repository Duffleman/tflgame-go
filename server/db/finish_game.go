package db

import (
	"context"
	"encoding/json"
	"time"

	"tflgame"

	"github.com/cuvva/ksuid-go"
)

// FinishGame ends a game. You **must** run this within a transaction.
func (qw *QueryableWrapper) FinishGame(ctx context.Context, userID, gameID string, score int, gameEndTime time.Time) error {
	eventID := ksuid.Generate("event").String()
	now := gameEndTime

	payloadBytes, err := json.Marshal(&tflgame.FinishGamePayload{
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
		eventID, "finish_game", userID, gameID, payloadBytes, now.Format(time.RFC3339),
	)
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
		UPDATE proj_games
		SET finished_at = $2, score = $3
		WHERE id = $1
	`, gameID, now.Format(time.RFC3339), score)

	return err
}
