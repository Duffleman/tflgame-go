package middleware

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"net/http"

	"tflgame/server/lib/cher"
	"tflgame/server/lib/httperr"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticateJWT(next http.HandlerFunc, publicKey *ecdsa.PublicKey) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("Authorization")

		if tokenString == "" {
			httperr.HandleError(res, cher.New(cher.Unauthorized, nil, cher.New("missing_header", nil)))
			return
		}

		var claims jwt.StandardClaims

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return publicKey, nil
		})
		if err != nil {
			httperr.HandleError(res, cher.New(cher.Unauthorized, nil, cher.Coerce(err)))
			return
		}

		if !token.Valid {
			httperr.HandleError(res, cher.New(cher.Unauthorized, nil, cher.New("invalid_token", nil)))
			return
		}

		newCtx := context.WithValue(req.Context(), TFLGameUser, claims.Subject)

		newReq := req.WithContext(newCtx)

		next.ServeHTTP(res, newReq)
	})
}

func AuthenticateInternalKey(next http.HandlerFunc, internalKeys []string) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("Authorization")

		if tokenString == "" {
			httperr.HandleError(res, cher.New(cher.Unauthorized, nil, cher.New("missing_header", nil)))
			return
		}

		keys := map[string]struct{}{}

		for _, k := range internalKeys {
			keys[k] = struct{}{}
		}

		if _, ok := keys[tokenString]; !ok {
			httperr.HandleError(res, cher.New(cher.Unauthorized, nil, cher.New("wrong_key", nil)))
			return
		}

		next.ServeHTTP(res, req)
	})
}
