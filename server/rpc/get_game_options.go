package rpc

import (
	"context"

	"tflgame"
)

func (r *RPC) GetGameOptions(ctx context.Context) (*tflgame.GetGameOptionsResponse, error) {
	return r.app.GetGameOptions(ctx)
}
