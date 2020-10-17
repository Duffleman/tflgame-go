package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"tflgame/server/lib/cher"
	"tflgame/server/lib/httperr"

	"github.com/xeipuuv/gojsonschema"
)

// ValidateSchema is middleware that allows for jsons schema validation
func ValidateSchema(next http.HandlerFunc, ls *gojsonschema.Schema) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if !strings.Contains(req.Header.Get("Content-Type"), "application/json") {
			httperr.HandleError(res, cher.New("unknown_content_type", nil))
			return
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			if netErr, ok := err.(net.Error); ok {
				httperr.HandleError(res, netErr)
				return
			}

			httperr.HandleError(res, fmt.Errorf("failed to read request body: %w", err))
			return
		}

		ld := gojsonschema.NewBytesLoader(body)

		result, err := ls.Validate(ld)
		if err != nil {
			httperr.HandleError(res, fmt.Errorf("schema validation failed: %w", err))
			return
		}

		err = coerceJSONSchemaError(result)
		if err != nil {
			httperr.HandleError(res, err)
			return
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body))

		next.ServeHTTP(res, req)
	})
}

func coerceJSONSchemaError(result *gojsonschema.Result) error {
	if result.Valid() {
		return nil
	}

	var reasons []cher.E

	errs := result.Errors()
	for _, err := range errs {
		reasons = append(reasons, cher.E{
			Code: "schema_failure",
			Meta: cher.M{
				"field":   err.Field(),
				"type":    err.Type(),
				"message": err.Description(),
			},
		})
	}

	return cher.New(cher.BadRequest, nil, reasons...)
}
