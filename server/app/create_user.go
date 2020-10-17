package app

import (
	"context"

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

	user, err := a.db.CreateUser(ctx, req.Handle, string(hash))
	if err != nil {
		return nil, err
	}

	token, err := a.GenerateJWT(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &tflgame.CreateUserResponse{
		ID:      user.ID,
		Handle:  user.Handle,
		Numeric: user.Numeric,
		Token:   token,
	}, nil
}
