package db

import (
	"context"
	"encoding/json"
	"time"

	"tflgame"

	"github.com/cuvva/ksuid-go"
)

func (d *DB) CreateGame(ctx context.Context, req *tflgame.CreateGameRequest, prompts []tflgame.Prompt) (string, error) {
	var gameID string

	err := d.DoTx(ctx, func(qw *QueryableWrapper) error {
		eventID := ksuid.Generate("event").String()
		gameID = ksuid.Generate("game").String()
		now := time.Now().Format(time.RFC3339)

		payloadBytes, err := json.Marshal(tflgame.CreateGamePayload{
			CreationID:        gameID,
			UserID:            req.UserID,
			Prompts:           prompts,
			DifficultyOptions: req.DifficultyOptions,
			GameOptions:       req.GameOptions,
		})
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
				INSERT INTO events
				(id, type, user_id, game_id, payload, created_at)
				VALUES($1, $2, $3, $4, $5, $6)
			`,
			eventID, "create_game", req.UserID, gameID, payloadBytes, now,
		)
		if err != nil {
			return err
		}

		linesBytes, err := json.Marshal(req.GameOptions.Lines)
		if err != nil {
			return err
		}

		zonesBytes, err := json.Marshal(req.GameOptions.Zones)
		if err != nil {
			return err
		}

		_, err = qw.q.ExecContext(ctx, `
				INSERT INTO proj_games
				(id, user_id, rounds, include_random_spaces, change_letter_order, reveal_word_length, lines, zones, score, created_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			`,
			gameID,
			req.UserID,
			req.DifficultyOptions.Rounds,
			req.DifficultyOptions.IncludeRandomSpaces,
			req.DifficultyOptions.ChangeLetterOrder,
			req.DifficultyOptions.RevealWordLength,
			linesBytes,
			zonesBytes,
			0,
			now,
		)
		if err != nil {
			return err
		}

		for _, p := range prompts {
			_, err = qw.q.ExecContext(ctx, `
					INSERT INTO proj_prompts
					(id, user_id, game_id, prompt, answer, correct, created_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7)
				`,
				p.ID,
				req.UserID,
				gameID,
				p.Prompt,
				p.Answer,
				false,
				now,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return gameID, err
}
