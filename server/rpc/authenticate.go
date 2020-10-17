package rpc

import (
	"context"
	"tflgame"

	"github.com/xeipuuv/gojsonschema"
)

var AuthenticateSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"handle",
		"numeric",
		"pin"
	],

	"properties": {
		"handle": {
			"type": "string",
			"pattern": "^[A-Z]{2,5}$"
		},

		"numeric": {
			"type": "string",
			"pattern": "^\\d{3}$"
		},

		"pin": {
			"type": ["null", "string"],
			"pattern": "^\\d{6}$"
		}
	}
}`)

func (r *RPC) Authenticate(ctx context.Context, req *tflgame.AuthenticateRequest) (*tflgame.AuthenticateResponse, error) {
	// unauthenticated requests allowed

	// TODO(gm): rate limit this
	return r.app.Authenticate(ctx, req)
}
