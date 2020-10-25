package rpc

import (
	"context"

	"tflgame"
)

func (r *RPC) GetLeaderboard(ctx context.Context) (*tflgame.GetLeaderboardResponse, error) {
	return r.app.GetLeaderboard(ctx)
}
