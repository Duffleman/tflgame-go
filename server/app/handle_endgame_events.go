package app

import (
	"context"

	"tflgame/server/db"
)

func (a *App) HandleEndgameEvents(ctx context.Context, can context.CancelFunc, gameID string) {
	defer can()

	err := a.db.DoTx(ctx, func(qw *db.QueryableWrapper) error {
		game, err := qw.GetGame(ctx, gameID)
		if err != nil {
			return err
		}

		prompts, err := a.db.Q.ListPrompts(ctx, gameID)
		if err != nil {
			return err
		}

		for _, p := range prompts {
			if p.AnsweredAt == nil {
				return nil
			}
		}

		score, _, err := a.CalculateGameScore(game.DifficultyOptions, prompts)
		if err != nil {
			return err
		}

		switch true {
		case game.FinishedAt == nil:
			err := qw.FinishGame(ctx, game.UserID, gameID, score)
			if err != nil {
				return err
			}
		case game.Score != score:
			err = qw.RecalculateGameScore(ctx, game.UserID, gameID, score)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		a.Logger.WithContext(ctx).WithError(err).Warn("cannot_calculate_player_score")
	}
}
