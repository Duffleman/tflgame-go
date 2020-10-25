package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var GetCurrentGameSchema = gojsonschema.NewStringLoader(`{
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

func (r *RPC) GetCurrentGame(ctx context.Context, req *tflgame.GetCurrentGameRequest) (*tflgame.GetCurrentGameResponse, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.GetCurrentGame(ctx, req)
}
