package app

import (
	"context"
	"fmt"
	"math"

	"tflgame"
	"tflgame/server/lib/ptr"
)

func (a *App) CalculateUserScore(ctx context.Context, userID string) (int, *tflgame.Calculations, error) {
	c := &tflgame.Calculations{}

	games, err := a.db.Q.ListAllGames(ctx, userID, true)
	if err != nil {
		return 0, nil, err
	}

	var score float64 = start

	c.Start = score
	c.Add(tflgame.CalculationEvent{
		Note:  "Start",
		Score: score,
	})

	var totalGameScore float64 = 0

	for _, g := range games {
		totalGameScore += float64(g.Score)

		c.Add(tflgame.CalculationEvent{
			Item:   &tflgame.CalculationItem{ID: g.ID, Type: "game"},
			Note:   "Game won",
			Effect: ptr.String(fmt.Sprintf("%d", g.Score)),
			Score:  totalGameScore,
		})
	}

	score = totalGameScore

	c.Base = score

	c.Add(tflgame.CalculationEvent{
		Note:  "Game score summed",
		Score: score,
	})

	score = score / float64(len(games))

	c.End = score

	c.Add(tflgame.CalculationEvent{
		Note:   "Game score averaged",
		Effect: ptr.String(fmt.Sprintf("/%d", len(games))),
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
