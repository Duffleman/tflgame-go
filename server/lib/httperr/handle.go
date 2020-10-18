package httperr

import (
	"encoding/json"
	"net/http"

	"tflgame/server/lib/cher"

	log "github.com/sirupsen/logrus"
)

// HandleError will handle returning errors to the writer for you
func HandleError(w http.ResponseWriter, e error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if v, ok := e.(cher.E); ok {
		switch v.Code {
		case cher.NotFound, cher.RouteNotFound:
			w.WriteHeader(404)
		default:
			w.WriteHeader(500)
		}
		json.NewEncoder(w).Encode(v)
		return
	}

	log.WithError(e).Warn("unknown")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": "unknown",
		"meta": map[string]interface{}{
			"error": e.Error(),
		},
	})

	return
}
