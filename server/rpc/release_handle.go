package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var ReleaseHandleSchema = gojsonschema.NewStringLoader(`{
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

func (r *RPC) ReleaseHandle(ctx context.Context, req *tflgame.ReleaseHandleRequest) error {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return cher.New(cher.Unauthorized, nil)
	}

	return r.app.ReleaseHandle(ctx, req)
}
