package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var ListEventsSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"user_id",
		"pagination"
	],

	"properties": {
		"user_id": {
			"type": "string",
			"minLength": 1
		},

		"pagination": {
			"type": ["object", "null"],
			"additionalProperties": false,

			"required": [
				"before",
				"after",
				"order",
				"limit"
			],

			"properties": {
				"before": {
					"type": ["string", "null"],
					"minLength": 1
				},

				"after": {
					"type": ["string", "null"],
					"minLength": 1
				},

				"order": {
					"type": "string",
					"enum": ["oldest_first", "newest_first"]
				},

				"limit": {
					"type": "integer",
					"minimum": 1,
					"maximum": 1000
				}
			}
		}
	}
}`)

func (r *RPC) ListEvents(ctx context.Context, req *tflgame.ListEventsRequest) ([]*tflgame.Event, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.ListEvents(ctx, req)
}
