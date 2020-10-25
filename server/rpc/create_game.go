package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/rpc/middleware"

	"github.com/xeipuuv/gojsonschema"
)

var CreateGameSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"user_id",
		"difficulty_options",
		"game_options"
	],

	"properties": {
		"user_id": {
			"type": "string",
			"minLength": 1
		},

		"difficulty_options": {
			"type": "object",
			"additionalProperties": false,

			"required": [
				"rounds",
				"include_random_spaces",
				"change_letter_order",
				"reveal_word_length"
			],

			"properties": {
				"rounds": {
					"type": "integer",
					"minimum": 10,
					"maximum": 100
				},

				"include_random_spaces": {
					"type": "boolean"
				},

				"change_letter_order": {
					"type": "boolean"
				},

				"reveal_word_length": {
					"type": "boolean"
				}
			}
		},

		"game_options": {
			"type": "object",
			"additionalProperties": false,

			"required": [
				"lines",
				"zones"
			],

			"properties": {
				"lines": {
					"type": "array",
					"minItems": 1,

					"items": {
						"type": "string",
						"minLength": 1
					}
				},

				"zones": {
					"type": ["null", "array"],
					"minItems": 1,

					"items": {
						"type": "string",
						"minLength": 1
					}
				}
			}
		}
	}
}`)

func (r *RPC) CreateGame(ctx context.Context, req *tflgame.CreateGameRequest) (*tflgame.CreateGameResponse, error) {
	userID := ctx.Value(middleware.TFLGameUser).(string)

	if req.UserID != userID {
		return nil, cher.New(cher.Unauthorized, nil)
	}

	return r.app.CreateGame(ctx, req)
}
