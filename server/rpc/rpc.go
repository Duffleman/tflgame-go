package rpc

import (
	"fmt"
	"net/http"

	"tflgame/server/app"
	"tflgame/server/lib/crpc"
	"tflgame/server/rpc/middleware"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

const (
	UnsafeNoAuth AuthType = iota
	JWT
)

type AuthType int

type RPC struct {
	app    *app.App
	router *chi.Mux
	Logger *logrus.Logger
}

func New(app *app.App, l *logrus.Logger) *RPC {
	return &RPC{
		app:    app,
		Logger: l,
		router: chi.NewRouter(),
	}
}

func (r *RPC) Route(pattern string, fnR interface{}, schema gojsonschema.JSONLoader, authRequirement AuthType) {
	fn := crpc.MustWrap(fnR)
	hasSchema := schema != nil

	if fn.AcceptsInput != hasSchema {
		if hasSchema {
			panic("schema validation configured, but handler doesn't accept input")
		} else {
			panic("no schema validation configured")
		}
	}

	var handler http.HandlerFunc = fn.Handler

	if schema != nil {
		compiledSchema, err := gojsonschema.NewSchemaLoader().Compile(schema)
		if err != nil {
			panic(fmt.Errorf("json schema error in %s: %w", pattern, err))
		}

		handler = middleware.ValidateSchema(handler, compiledSchema)
	}

	switch authRequirement {
	case UnsafeNoAuth:
	case JWT:
		handler = middleware.Authenticate(handler, r.app.SigningKeys.GetPublicKey())
	}

	r.router.Post(pattern, handler)
}

// Serve starts listening on the address for web traffic
func (r *RPC) Serve(address string) {
	r.Logger.Infof("starting web server on %s", address)
	r.Logger.Fatal(http.ListenAndServe(address, r.router))
}
