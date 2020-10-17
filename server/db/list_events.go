package db

import (
	"context"

	"tflgame"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) ListEvents(ctx context.Context, req *tflgame.ListEventsRequest) ([]*tflgame.Event, error) {
	q := NewQueryBuilder().
		Select("id", "type", "user_id", "game_id", "payload", "created_at").
		From("events e").
		Where(sq.Eq{
			"user_id": req.UserID,
		})

	p, err := NewPagination(req.Pagination)
	if err != nil {
		return nil, err
	}

	if p.before != nil {
		q = q.Where("e.id < ?", p.before)
	}

	if p.after != nil {
		q = q.Where("e.id > ?", p.after)
	}

	q = q.
		OrderBy("e.serial " + p.order).
		Limit(uint64(p.limit))

	query, values, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []*tflgame.Event{}

	for rows.Next() {
		var e tflgame.Event
		var payloadBytes []byte

		err := rows.Scan(&e.ID, &e.Type, &e.UserID, &e.GameID, &payloadBytes, &e.CreatedAt)
		if err != nil {
			return nil, err
		}

		if err := tflgame.PayloadHandler(e.Type, payloadBytes, &e.Payload); err != nil {
			return nil, err
		}

		safePayload, err := tflgame.SafePublicPayload(e.Payload)
		if err != nil {
			return nil, err
		}

		e.Payload = safePayload

		events = append(events, &e)
	}

	return events, nil
}
