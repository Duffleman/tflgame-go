package rpc

import (
	"context"

	"tflgame"

	"github.com/xeipuuv/gojsonschema"
)

var TestGameOptionsSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"modes",
		"lines",
		"zones"
	],

	"properties": {
		"modes": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			}
		},

		"lines": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			}
		},

		"zones": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			}
		}
	}
}`)

func (r *RPC) TestGameOptions(ctx context.Context, req *tflgame.GameOptions) (*tflgame.TestGameOptionsResponse, error) {
	return r.app.TestGameOptions(ctx, req)
}
