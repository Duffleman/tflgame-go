package db

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) ListAllGames(ctx context.Context, userID string, onlyFinished bool) ([]*DBGame, error) {
	qb := NewQueryBuilder().
		Select("e.type, e.payload, e.user_id, g.created_at, g.score, g.finished_at").
		From("events e").
		Join("proj_games g ON e.game_id = g.id").
		Where(sq.Eq{
			"e.user_id": userID,
			"e.type":    "create_game",
		})

	if onlyFinished {
		qb = qb.Where(sq.NotEq{
			"g.finished_at": nil,
		})
	}

	query, values, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := []*DBGame{}

	for rows.Next() {
		var eventType string
		var payloadBytes []byte
		var g DBGame

		err := rows.Scan(&eventType, &payloadBytes, &g.UserID, &g.CreatedAt, &g.Score, &g.FinishedAt)
		if err != nil {
			return nil, err
		}

		var i interface{}

		err = tflgame.PayloadHandler(eventType, payloadBytes, &i)
		if err != nil {
			return nil, err
		}

		payload, ok := i.(tflgame.CreateGamePayload)
		if !ok {
			return nil, cher.New("payload_mismatch", nil)
		}

		g.DifficultyOptions = payload.DifficultyOptions
		g.GameOptions = payload.GameOptions

		set = append(set, &g)
	}

	return set, nil
}
