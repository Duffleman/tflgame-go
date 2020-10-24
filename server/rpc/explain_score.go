package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var ExplainScoreSchema = gojsonschema.NewStringLoader(`{
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
			"type": ["string", "null"],
			"minLength": 1
		}
	}
}`)

func (r *RPC) ExplainScore(ctx context.Context, req *tflgame.ExplainScoreRequest) (*tflgame.Calculations, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.ExplainScore(ctx, req)
}
