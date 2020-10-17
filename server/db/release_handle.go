package db

import (
	"context"
	"encoding/json"
	"time"

	ksuid "github.com/cuvva/ksuid-go"
)

func (d *DB) ReleaseHandle(ctx context.Context, userID string) error {
	return d.DoTx(ctx, func(qw *QueryableWrapper) error {
		user, err := qw.GetUserByID(ctx, userID)
		if err != nil {
			return err
		}

		eventID := ksuid.Generate("event").String()

		payloadBytes, err := json.Marshal(map[string]interface{}{
			"user_id": userID,
			"handle":  user.Handle,
			"numeric": user.Numeric,
		})
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
				INSERT INTO events
				(id, type, user_id, game_id, payload, created_at)
				VALUES($1, $2, $3, $4, $5, $6)
			`,
			eventID, "release_handle", userID, nil, payloadBytes, time.Now(),
		)
		if err != nil {
			return err
		}

		// TODO(gm): this won't work when games are attached
		// will need to release all games first
		_, err = qw.q.ExecContext(ctx, `
			DELETE FROM proj_users
			WHERE id = $1
		`, userID)

		return err
	})
}