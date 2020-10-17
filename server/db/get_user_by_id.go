package db

import (
	"context"

	"tflgame"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) GetUserByID(ctx context.Context, userID string) (*tflgame.User, error) {
	qb := NewQueryBuilder()
	query, values, err := qb.
		Select("id, handle, numeric, score, pin, created_at").
		From("proj_users u").
		Where(sq.Eq{
			"u.id": userID,
		}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := qw.q.QueryRowContext(ctx, query, values...)

	var u tflgame.User

	if err := row.Scan(&u.ID, &u.Handle, &u.Numeric, &u.Score, &u.Pin, &u.CreatedAt); err != nil {
		return nil, coerceNotFound(err)
	}

	return &u, nil
}
