package db

import (
	"context"
	"time"

	"tflgame"
	"tflgame/server/lib/cher"

	sq "github.com/Masterminds/squirrel"
)

type DBGame struct {
	UserID            string
	Score             int
	CreatedAt         time.Time
	FinishedAt        *time.Time
	DifficultyOptions tflgame.DifficultyOptions
	GameOptions       tflgame.GameOptions
}

func (qw *QueryableWrapper) GetGame(ctx context.Context, gameID string) (*DBGame, error) {
	query, values, err := NewQueryBuilder().
		Select("e.type, e.payload, e.user_id, g.created_at, g.score, g.finished_at").
		From("events e").
		Join("proj_games g ON e.game_id = g.id").
		Where(sq.Eq{
			"game_id": gameID,
			"type":    "create_game",
		}).
		Limit(uint64(1)).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := qw.q.QueryRowContext(ctx, query, values...)

	var g DBGame
	var eventType string
	var payloadBytes []byte

	if err := row.Scan(&eventType, &payloadBytes, &g.UserID, &g.CreatedAt, &g.Score, &g.FinishedAt); err != nil {
		return nil, coerceNotFound(err)
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

	return &g, nil
}
