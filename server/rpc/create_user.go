package rpc

import (
	"context"

	"tflgame"

	"github.com/xeipuuv/gojsonschema"
)

var CreateUserSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"additionalProperties": false,

	"required": [
		"handle",
		"pin"
	],

	"properties": {
		"handle": {
			"type": "string",
			"pattern": "^[A-Z]{2,5}$"
		},

		"pin": {
			"type": ["null", "string"],
			"pattern": "^\\d{6}$"
		}
	}
}`)

func (r *RPC) CreateUser(ctx context.Context, req *tflgame.CreateUserRequest) (*tflgame.CreateUserResponse, error) {
	// unauthenticated requests allowed

	// TODO(gm): rate limit this
	return r.app.CreateUser(ctx, req)
}
