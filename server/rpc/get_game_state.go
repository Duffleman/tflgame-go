package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var GetGameStateSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"user_id",
		"game_id"
	],

	"properties": {
		"user_id": {
			"type": "string",
			"minLength": 1
		},

		"game_id": {
			"type": "string",
			"minLength": 1
		}
	}
}`)

func (r *RPC) GetGameState(ctx context.Context, req *tflgame.GetGameStateRequest) (*tflgame.GetGameStateResponse, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.GetGameState(ctx, req)
}
