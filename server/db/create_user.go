package db

import (
	"context"
	"encoding/json"
	"tflgame"
	"time"

	ksuid "github.com/cuvva/ksuid-go"
)

func (d *DB) CreateUser(ctx context.Context, handle, hash string) (*tflgame.PublicUser, error) {
	user := &tflgame.PublicUser{}

	var insertHash *string = nil

	if hash != "" {
		insertHash = &hash
	}

	err := d.DoTx(ctx, func(qw *QueryableWrapper) error {
		numeric, err := qw.getMaxNumeric(ctx, handle)
		if err != nil {
			return err
		}

		eventID := ksuid.Generate("event").String()
		userID := ksuid.Generate("user").String()
		now := time.Now()

		payloadBytes, err := json.Marshal(map[string]interface{}{
			"creation_id": userID,
			"handle":      handle,
			"numeric":     numeric,
			"pin":         insertHash,
		})
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
				INSERT INTO events
				(id, type, user_id, game_id, payload, created_at)
				VALUES($1, $2, $3, $4, $5, $6)
			`,
			eventID, "create_user", userID, nil, payloadBytes, now,
		)
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
			INSERT INTO proj_users
			(id, handle, numeric, pin, score, created_at)
			VALUES($1, $2, $3, $4, 0, $5)
		`, userID, handle, numeric, insertHash, now)

		user.Handle = handle
		user.Numeric = numeric
		user.ID = userID

		return err
	})

	return user, err
}
