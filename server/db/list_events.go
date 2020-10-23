package db

import (
	"context"

	"tflgame"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) ListEvents(ctx context.Context, req *tflgame.ListEventsRequest) ([]*tflgame.Event, error) {
	q := NewQueryBuilder().
		Select("e.id", "e.type", "e.user_id", "e.game_id", "e.payload", "e.created_at").
		From("events e").
		LeftJoin("proj_games g ON e.game_id = g.id").
		Where(sq.Eq{
			"e.user_id": req.UserID,
		}).
		// only finished games appear in the timeline
		Where(sq.Or{
			sq.NotEq{
				"g.id":          nil,
				"g.finished_at": nil,
			},
			sq.Eq{
				"g.id": nil,
			},
		})

	p, err := NewPagination(req.Pagination)
	if err != nil {
		return nil, err
	}

	if p.before != nil {
		q = q.Where("e.serial < (SELECT serial FROM events WHERE id = ?)", p.before)
	}

	if p.after != nil {
		q = q.Where("e.serial > (SELECT serial FROM events WHERE id = ?)", p.after)
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
