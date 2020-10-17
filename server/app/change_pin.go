package app

import (
	"context"
	"tflgame/server/lib/cher"

	"tflgame"
)

func (a *App) ChangePin(ctx context.Context, req *tflgame.ChangePinRequest) error {
	user, err := a.db.Q.GetUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if user.Pin != nil {
		return cher.New(cher.Unauthorized, nil)
	}

	if err := a.CheckPin(user, &req.CurrentPin); err != nil {
		return err
	}

	return nil
}
