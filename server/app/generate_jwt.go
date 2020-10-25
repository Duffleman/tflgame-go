package app

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (a *App) GenerateJWT(ctx context.Context, userID string, expiry time.Duration) (string, error) {
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &jwt.StandardClaims{
		Audience:  "web",
		ExpiresAt: now.Add(expiry).Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    "auth.tflgame.com",
		NotBefore: now.Unix(),
		Subject:   userID,
	})

	return a.SigningKeys.SignToken(token)
}
