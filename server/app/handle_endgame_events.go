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

		gameScore, _, err := a.CalculateGameScore(game.DifficultyOptions, prompts)
		if err != nil {
			return err
		}

		switch true {
		case game.FinishedAt == nil:
			err := qw.FinishGame(ctx, game.UserID, gameID, gameScore)
			if err != nil {
				return err
			}
		case game.Score != gameScore:
			err = qw.RecalculateGameScore(ctx, game.UserID, gameID, gameScore)
			if err != nil {
				return err
			}
		}

		// TODO(gm): parallelise if slow
		user, err := qw.GetUserByID(ctx, game.UserID)
		if err != nil {
			return err
		}

		userScore, c, err := a.CalculateUserScore(ctx, game.UserID)
		if err != nil {
			return err
		}

		if user.Score != userScore {
			qw.RecalculateUserScore(ctx, user.ID, userScore)
		}

		return nil
	})
	if err != nil {
		a.Logger.WithContext(ctx).WithError(err).Warn("cannot_calculate_player_score")
	}
}
