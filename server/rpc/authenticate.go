package rpc

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
	"tflgame/server/lib/crpc"

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

	rc := crpc.GetRequestContext(ctx)
	ip := stripPort(rc.RemoteAddr)

	allowed := r.limiter.Allow(ip)
	if !allowed {
		return nil, cher.New(cher.TooManyRequests, nil)
	}

	return r.app.Authenticate(ctx, req)
}
