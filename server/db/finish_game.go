package db

import (
	"context"
	"time"

	"github.com/cuvva/ksuid-go"
)

func (qw *QueryableWrapper) FinishGame(ctx context.Context, userID, gameID string) error {
	eventID := ksuid.Generate("event").String()
	now := time.Now()

	_, err := qw.q.ExecContext(ctx, `
			INSERT INTO events
			(id, type, user_id, game_id, created_at)
			VALUES($1, $2, $3, $4, $5)
		`,
		eventID, "finish_game", userID, gameID, now.Format(time.RFC3339),
	)
	if err != nil {
		return err
	}

	_, err = qw.q.ExecContext(ctx, `
		UPDATE proj_games
		SET finished_at = $2
		WHERE id = $1
	`, gameID, now.Format(time.RFC3339))

	return err
}
