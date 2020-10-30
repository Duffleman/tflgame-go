package middleware

import (
	"net/http"
)

func AddCORSHeaders(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(res, req)
	})
}
