package app

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"

	"golang.org/x/crypto/bcrypt"
)

func (a *App) ChangePin(ctx context.Context, req *tflgame.ChangePinRequest) error {
	user, err := a.db.Q.GetUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if user.Pin == nil {
		return cher.New(cher.Unauthorized, nil)
	}

	if err := a.CheckPin(user, &req.CurrentPin); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPin), bcrypt.MinCost)
	if err != nil {
		return err
	}

	return a.db.ChangePin(ctx, user.ID, hash)
}
