package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"tflgame"

	sq "github.com/Masterminds/squirrel"
	ksuid "github.com/cuvva/ksuid-go"
)

func (d *DB) ChangeHandle(ctx context.Context, userID, handle string) (string, error) {
	var newNumeric string

	err := d.DoTx(ctx, func(qw *QueryableWrapper) error {
		numeric, err := qw.getMaxNumeric(ctx, handle)
		if err != nil {
			return err
		}

		newNumeric = numeric

		eventID := ksuid.Generate("event").String()

		payloadBytes, err := json.Marshal(tflgame.ChangeHandlePayload{
			UserID:     userID,
			NewHandle:  handle,
			NewNumeric: numeric,
		})
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
				INSERT INTO events
				(id, type, user_id, game_id, payload, created_at)
				VALUES($1, $2, $3, $4, $5, $6)
			`,
			eventID, "change_handle", userID, nil, payloadBytes, time.Now().Format(time.RFC3339),
		)
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
			UPDATE proj_users
			SET handle = $2, numeric = $3
			WHERE id = $1
		`, userID, handle, numeric)

		return err
	})

	return newNumeric, err
}

func (qw *QueryableWrapper) getMaxNumeric(ctx context.Context, handle string) (string, error) {
	qb := NewQueryBuilder()
	query, values, err := qb.
		Select("MAX(numeric)").
		From("proj_users u").
		Where(sq.Eq{
			"u.handle": handle,
		}).
		Limit(1).
		ToSql()
	if err != nil {
		return "", err
	}

	row := qw.q.QueryRowContext(ctx, query, values...)

	var numeric string
	var decode *string

	if err := row.Scan(&decode); err != nil {
		if err == sql.ErrNoRows {
			numeric = "000"
		}

		return "", err
	}

	if decode != nil {
		numeric = *decode
	} else {
		numeric = "000"
	}

	numberForm, err := strconv.Atoi(numeric)
	if err != nil {
		return "", err
	}

	numberForm++
	numeric = fmt.Sprintf("%03s", strconv.Itoa(numberForm))

	return numeric, nil
}
