package db

import (
	"context"
)

func (qw *QueryableWrapper) GetCurrentGame(ctx context.Context, userID string) (string, error) {
	var existingGameID string
	row := qw.q.QueryRowContext(ctx, `
		SELECT id
		FROM proj_games
		WHERE
			user_id = $1
			AND finished_at IS NULL
	`, userID)

	err := row.Scan(&existingGameID)
	if err != nil {
		return "", err
	}

	return existingGameID, nil
}
