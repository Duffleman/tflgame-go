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
	CreatedAt         time.Time
	FinishedAt        *time.Time
	DifficultyOptions tflgame.DifficultyOptions
	GameOptions       tflgame.GameOptions
}

func (qw *QueryableWrapper) GetGameOpts(ctx context.Context, gameID string) (*DBGame, error) {
	query, values, err := NewQueryBuilder().
		Select("e.type, e.payload, e.user_id, g.created_at, g.finished_at").
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

	var userID string
	var createdAt time.Time
	var finishedAt *time.Time
	var eventType string
	var payloadBytes []byte

	if err := row.Scan(&eventType, &payloadBytes, &userID, &createdAt, &finishedAt); err != nil {
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

	return &DBGame{
		UserID:            userID,
		CreatedAt:         createdAt,
		FinishedAt:        finishedAt,
		DifficultyOptions: payload.DifficultyOptions,
		GameOptions:       payload.GameOptions,
	}, nil
}
