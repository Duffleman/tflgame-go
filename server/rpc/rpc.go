package rpc

import (
	"fmt"
	"net/http"

	"tflgame/server/app"
	"tflgame/server/lib/cher"
	"tflgame/server/lib/crpc"
	"tflgame/server/lib/httperr"
	"tflgame/server/rpc/middleware"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

const (
	UnsafeNoAuth AuthType = iota
	InternalOnlyAuth
	JWTAuth
)

type AuthType int

type RPC struct {
	app         *app.App
	router      *chi.Mux
	internalKey string
	Logger      *logrus.Logger
}

func New(app *app.App, l *logrus.Logger, internalKey string) *RPC {
	r := &RPC{
		app:         app,
		Logger:      l,
		internalKey: internalKey,
		router:      chi.NewRouter(),
	}

	r.SetMuxBase()

	return r
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
	case InternalOnlyAuth:
		handler = middleware.AuthenticateInternalKey(handler, r.internalKey)
	case JWTAuth:
		handler = middleware.AuthenticateJWT(handler, r.app.SigningKeys.GetPublicKey())
	}

	r.router.Post(pattern, handler)
}

func (r *RPC) SetMuxBase() {
	r.router.NotFound(func(res http.ResponseWriter, req *http.Request) {
		httperr.HandleError(res, cher.New(cher.RouteNotFound, nil))
		return
	})
}

// Serve starts listening on the address for web traffic
func (r *RPC) Serve(address string) {
	r.Logger.Infof("starting web server on %s", address)
	r.Logger.Fatal(http.ListenAndServe(address, r.router))
}
