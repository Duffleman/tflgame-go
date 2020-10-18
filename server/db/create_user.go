package db

import (
	"context"
	"encoding/json"
	"tflgame"
	"time"

	ksuid "github.com/cuvva/ksuid-go"
)

func (d *DB) CreateUser(ctx context.Context, handle string, hash []byte) (*tflgame.PublicUser, error) {
	user := &tflgame.PublicUser{}

	var insertHash *string

	if len(hash) > 0 {
		ih := string(hash)
		insertHash = &ih
	}

	err := d.DoTx(ctx, func(qw *QueryableWrapper) error {
		numeric, err := qw.getMaxNumeric(ctx, handle)
		if err != nil {
			return err
		}

		eventID := ksuid.Generate("event").String()
		userID := ksuid.Generate("user").String()
		now := time.Now().Format(time.RFC3339)

		payloadBytes, err := json.Marshal(tflgame.CreateUserPayload{
			CreationID: userID,
			Handle:     handle,
			Numeric:    numeric,
			Pin:        insertHash,
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
		user.UserID = userID

		return err
	})

	return user, err
}
