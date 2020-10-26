package crpc

import (
	"context"
	"net/http"
)

type contextKey string

const requestKey contextKey = "crpcrequest"

// GetRequestContext returns the Request from the context object
func GetRequestContext(ctx context.Context) *http.Request {
	if val, ok := ctx.Value(requestKey).(*http.Request); ok {
		return val
	}

	return nil
}

func setRequestContext(ctx context.Context, request *http.Request) context.Context {
	return context.WithValue(ctx, requestKey, request)
}
