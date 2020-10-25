package app

import (
	"context"
	"fmt"
	"math"

	"tflgame"
	"tflgame/server/db"
	"tflgame/server/lib/ptr"
)

func (a *App) CalculateUserScore(ctx context.Context, qw *db.QueryableWrapper, userID string) (int, *tflgame.Calculations, error) {
	if qw == nil {
		qw = a.db.Q
	}

	c := &tflgame.Calculations{}

	games, err := qw.ListAllGames(ctx, userID, true)
	if err != nil {
		return 0, nil, err
	}

	var score float64 = start

	c.Start = score

	c.Add(tflgame.CalculationEvent{
		Note:  "Start",
		Score: score,
	})

	for _, g := range games {
		score += float64(g.Score)

		c.Add(tflgame.CalculationEvent{
			Item:   &tflgame.CalculationItem{ID: g.ID, Type: "game"},
			Note:   "Game won",
			Effect: ptr.String(fmt.Sprintf("%d", g.Score)),
			Score:  score,
		})
	}

	c.Base = score

	c.Add(tflgame.CalculationEvent{
		Note:  "Game score summed",
		Score: score,
	})

	gamesLength := 1

	if len(games) > 0 {
		gamesLength = len(games)
	}

	score = score / float64(gamesLength)

	c.End = score

	c.Add(tflgame.CalculationEvent{
		Note:   "Game score averaged",
		Effect: ptr.String(fmt.Sprintf("/%d", gamesLength)),
		Score:  score,
	})

	final := int(math.Ceil(score))

	c.Final = final

	c.Add(tflgame.CalculationEvent{
		Note:  "Score rounded up",
		Score: float64(final),
	})

	return final, c, nil
}
