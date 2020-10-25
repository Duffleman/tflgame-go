package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/lib/crpc"

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

	rc := crpc.GetRequestContext(ctx)
	ip := stripPort(rc.RemoteAddr)

	allowed := r.limiter.Allow(ip)
	if !allowed {
		return nil, cher.New(cher.TooManyRequests, nil)
	}

	return r.app.CreateUser(ctx, req)
}
