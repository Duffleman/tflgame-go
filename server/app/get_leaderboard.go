package app

import (
	"context"

	"tflgame"
)

func (a *App) GetLeaderboard(ctx context.Context) (*tflgame.GetLeaderboardResponse, error) {
	users, err := a.db.Q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var players []*tflgame.Player

	for _, user := range users {
		players = append(players, &tflgame.Player{
			ID:      user.ID,
			Handle:  user.Handle,
			Numeric: user.Numeric,
			Score:   user.Score,
		})
	}

	return &tflgame.GetLeaderboardResponse{
		Level:   "All",
		Color:   "#FFFFFF",
		Players: players,
	}, nil
}
