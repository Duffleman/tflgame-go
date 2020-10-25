package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var ListGameHistorySchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"user_id"
	],

	"properties": {
		"user_id": {
			"type": "string",
			"minLength": 1
		}
	}
}`)

func (r *RPC) ListGameHistory(ctx context.Context, req *tflgame.ListGameHistoryRequest) ([]*tflgame.GameLog, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.ListGameHistory(ctx, req)
}
