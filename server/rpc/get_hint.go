package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var GetHintSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"user_id",
		"prompt_id"
	],

	"properties": {
		"user_id": {
			"type": "string",
			"minLength": 1
		},

		"prompt_id": {
			"type": "string",
			"minLength": 1
		}
	}
}`)

func (r *RPC) GetHint(ctx context.Context, req *tflgame.GetHintRequest) (*tflgame.GetHintResponse, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.GetHint(ctx, req)
}
