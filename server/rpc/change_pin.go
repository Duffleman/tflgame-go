package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var ChangePinSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"user_id",
		"current_pin",
		"new_pin"
	],

	"properties": {
		"user_id": {
			"type": "string",
			"minLength": 1
		},

		"current_pin": {
			"type": ["null", "string"],
			"pattern": "^\\d{6}$"
		},

		"new_pin": {
			"type": ["null", "string"],
			"pattern": "^\\d{6}$"
		}
	}
}`)

func (r *RPC) ChangePin(ctx context.Context, req *tflgame.ChangePinRequest) error {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return cher.New(cher.Unauthorized, nil)
	}

	return r.app.ChangePin(ctx, req)
}
