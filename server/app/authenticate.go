package app

import (
	"context"
	"time"

	"tflgame"
	"tflgame/server/lib/cher"
)

func (a *App) Authenticate(ctx context.Context, req *tflgame.AuthenticateRequest) (*tflgame.AuthenticateResponse, error) {
	user, err := a.db.Q.GetUserByTag(ctx, req.Handle, req.Numeric)
	if err != nil {
		return nil, err
	}

	if req.Pin == "" {
		return nil, cher.New("disallowed", nil)
	}

	err = a.CheckPin(user, &req.Pin)
	if err != nil {
		return nil, err
	}

	token, err := a.GenerateJWT(ctx, user.ID, 8766*time.Hour)
	if err != nil {
		return nil, err
	}

	return &tflgame.AuthenticateResponse{
		UserID: user.ID,
		Token:  token,
	}, nil
}
