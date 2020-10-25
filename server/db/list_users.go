package db

import (
	"context"

	"tflgame"
)

func (qw *QueryableWrapper) ListUsers(ctx context.Context) ([]*tflgame.User, error) {
	query, values, err := NewQueryBuilder().
		Select("id, handle, numeric, score, pin, created_at").
		From("proj_users u").
		OrderBy("score ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*tflgame.User{}

	for rows.Next() {
		var u tflgame.User

		if err := rows.Scan(&u.ID, &u.Handle, &u.Numeric, &u.Score, &u.Pin, &u.CreatedAt); err != nil {
			return nil, coerceNotFound(err)
		}

		users = append(users, &u)
	}

	return users, nil
}
