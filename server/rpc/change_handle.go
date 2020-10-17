package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var ChangeHandleSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"id",
		"new_handle"
	],

	"properties": {
		"id": {
			"type": "string",
			"minLength": 1
		},

		"new_handle": {
			"type": "string",
			"minLength": 2,
			"maxLength": 5
		}
	}
}`)

func (r *RPC) ChangeHandle(ctx context.Context, req *tflgame.ChangeHandleRequest) (*tflgame.PublicUser, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.ID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.ChangeHandle(ctx, req)
}
