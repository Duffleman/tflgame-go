package app

import (
	"context"
	"time"

	"tflgame"

	"golang.org/x/crypto/bcrypt"
)

func (a *App) CreateUser(ctx context.Context, req *tflgame.CreateUserRequest) (*tflgame.CreateUserResponse, error) {
	var hash []byte

	if req.Pin != nil {
		bhash, err := bcrypt.GenerateFromPassword([]byte(*req.Pin), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}

		hash = bhash
	}

	user, err := a.db.CreateUser(ctx, req.Handle, hash)
	if err != nil {
		return nil, err
	}

	var jwtExpiry time.Duration

	if req.Pin == nil {
		jwtExpiry = 24 * time.Hour // 1 day
	} else {
		jwtExpiry = 8766 * time.Hour // 1 year
	}

	token, err := a.GenerateJWT(ctx, user.UserID, jwtExpiry)
	if err != nil {
		return nil, err
	}

	return &tflgame.CreateUserResponse{
		ID:      user.UserID,
		Handle:  user.Handle,
		Numeric: user.Numeric,
		Token:   token,
	}, nil
}
