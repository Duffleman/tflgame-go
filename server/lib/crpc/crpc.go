package crpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"tflgame/server/lib/cher"
	"tflgame/server/lib/httperr"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()
var contextType = reflect.TypeOf((*context.Context)(nil)).Elem()

type WrappedFunc struct {
	Handler       http.HandlerFunc
	AcceptsInput  bool
	ReturnsResult bool
}

// Wrap reflects a HandlerFunc from any function matching the
// following signatures:
//
// func(ctx context.Context, request *T) (response *T, err error)
// func(ctx context.Context, request *T) (err error)
// func(ctx context.Context) (response *T, err error)
// func(ctx context.Context) (err error)
func Wrap(fn interface{}) (*WrappedFunc, error) {
	// prevent re-reflection of type that is already a HandlerFunc
	if _, ok := fn.(http.HandlerFunc); ok {
		return nil, fmt.Errorf("fn doesn't need to be wrapped, use RegisterFunc")
	}

	v := reflect.ValueOf(fn)
	t := v.Type()

	// check the basic type and the number of inputs/outputs
	if t.Kind() != reflect.Func {
		return nil, fmt.Errorf("fn must be function, got %s", t.Kind())
	} else if t.NumIn() < 1 || t.NumIn() > 2 {
		return nil, fmt.Errorf("fn input must be (context.Context) or (context.Context, *T), got %d arguments", t.NumIn())
	} else if t.NumOut() < 1 || t.NumOut() > 2 {
		return nil, fmt.Errorf("fn output must be (error) or (*T, error), got %d arguments", t.NumOut())
	}

	if !t.In(0).Implements(contextType) {
		return nil, fmt.Errorf("fn first argument must implement context.Context, got %s", t.In(0))
	} else if !t.Out(t.NumOut() - 1).Implements(errorType) {
		return nil, fmt.Errorf("fn last argument must implement error, got %s", t.Out(t.NumOut()-1))
	}

	var reqT, resT reflect.Type = nil, nil

	if t.NumIn() == 2 {
		if t.In(1).Kind() != reflect.Ptr {
			return nil, fmt.Errorf("fn last argument must be a pointer, got %s", t.In(1))
		}

		reqT = t.In(1).Elem()
		if reqT.Kind() != reflect.Struct {
			return nil, fmt.Errorf("fn last argument must be a struct, got %s", reqT.Kind())
		}
	}

	if t.NumOut() == 2 {
		var err error

		resT, err = wrapReturn(t.Out(0))
		if err != nil {
			return nil, err
		}
	}

	hn := func(w http.ResponseWriter, r *http.Request) {
		ctx := reflect.ValueOf(r.Context())
		var inputs []reflect.Value

		if reqT == nil {
			if r.Body != nil {
				i, err := r.Body.Read(make([]byte, 1))
				if i != 0 || err != io.EOF {
					httperr.HandleError(w, cher.New(cher.BadRequest, nil, cher.New("unexpected_request_body", nil)))
					return
				}
			}

			inputs = []reflect.Value{ctx}
		} else {
			if r.Body == nil {
				httperr.HandleError(w, cher.New(cher.BadRequest, nil, cher.New("missing_request_body", nil)))
				return
			}

			req := reflect.New(reqT)
			err := json.NewDecoder(r.Body).Decode(req.Interface())
			if err == io.EOF {
				httperr.HandleError(w, cher.New(cher.BadRequest, nil, cher.New("missing_request_body", nil)))
				return
			} else if err != nil {
				httperr.HandleError(w, fmt.Errorf("crpc: json decoder error: %w", err))
				return
			}

			inputs = []reflect.Value{ctx, req}
		}

		res := v.Call(inputs)

		if err := res[len(res)-1]; !err.IsNil() {
			httperr.HandleError(w, err.Interface().(error))
			return
		}

		if len(res) == 1 {
			w.WriteHeader(http.StatusNoContent)
		} else if len(res) == 2 {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(res[0].Interface())
			if err != nil {
				httperr.HandleError(w, err)
				return
			}
		}

		return
	}

	return &WrappedFunc{
		Handler:       hn,
		AcceptsInput:  reqT != nil,
		ReturnsResult: resT != nil,
	}, nil
}

func wrapReturn(t reflect.Type) (reflect.Type, error) {
	switch t.Kind() {
	case reflect.Ptr:
		n := t.Elem()

		if n.Kind() != reflect.Struct {
			_, err := wrapReturn(n)
			if err != nil {
				return nil, err
			}
		}

		return n, nil

	case reflect.Slice:
		n := t.Elem()

		if n.Kind() != reflect.String {
			_, err := wrapReturn(n)
			if err != nil {
				return nil, err
			}
		}

		return t, nil

	default:
		return nil, fmt.Errorf("unsupported return type, expected *struct or slice; got %s", t)
	}
}

// MustWrap is the same as Wrap, however it panics when passed an
// invalid handler.
func MustWrap(fn interface{}) *WrappedFunc {
	wrapped, err := Wrap(fn)
	if err != nil {
		panic(err)
	}

	return wrapped
}
