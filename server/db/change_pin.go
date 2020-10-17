package db

import (
	"context"
	"encoding/json"
	"time"

	ksuid "github.com/cuvva/ksuid-go"
)

func (d *DB) ChangePin(ctx context.Context, userID string, hash []byte) error {
	return d.DoTx(ctx, func(qw *QueryableWrapper) error {
		eventID := ksuid.Generate("event").String()

		payloadBytes, err := json.Marshal(map[string]interface{}{
			"user_id": userID,
			"pin":     hash,
		})
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
				INSERT INTO events
				(id, type, user_id, game_id, payload, created_at)
				VALUES($1, $2, $3, $4, $5, $6)
			`,
			eventID, "change_pin", userID, nil, payloadBytes, time.Now(),
		)
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
			UPDATE proj_users
			SET pin = $2
			WHERE id = $1
		`, userID, hash)

		return err
	})
}
