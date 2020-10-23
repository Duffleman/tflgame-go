package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var SubmitAnswerSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"user_id",
		"prompt_id",
		"answer"
	],

	"properties": {
		"user_id": {
			"type": "string",
			"minLength": 1
		},

		"prompt_id": {
			"type": "string",
			"minLength": 1
		},

		"answer": {
			"type": "string",
			"minLength": 1
		}
	}
}`)

func (r *RPC) SubmitAnswer(ctx context.Context, req *tflgame.SubmitAnswerRequest) (*tflgame.SubmitAnswerResponse, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.SubmitAnswer(ctx, req)
}
